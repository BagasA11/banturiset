package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strings"

	xendit "github.com/xendit/xendit-go/v5"
	invoice "github.com/xendit/xendit-go/v5/invoice"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
)

type DonasiService struct {
	Repo *repository.DonasiRepo
}

// type fee struct {
// 	Typ   string  `json:"type"`
// 	Value float32 `json:"value"`
// }

// const paymenturl = "https://api.xendit.co/v2/invoices"

const getInvoiceUrl = "https://api.xendit.co/v2/invoices/"

func NewDonasiService() *DonasiService {
	return &DonasiService{
		Repo: repository.NewDonasiRepo(),
	}
}

func (ds *DonasiService) Create(donaturID uint, ProjectID uint, req dto.CreateDonasi) (*models.Donasi, error) {
	if !slices.Contains([]string{"ovo", "bca", "bsi", "mandiri", "bri", "bni", "bjb", "shopeepay"}, strings.ToLower(req.Method)) {
		return nil, fmt.Errorf("metode pembayaran %s tidak didukung", req.Method)
	}

	admin := req.Jml * 0.05
	var fee float32

	if slices.Contains([]string{"bca", "bsi", "mandiri"}, strings.ToLower(req.Method)) {
		fee = float32(6000) + admin
	}

	if strings.ToLower(req.Method) == "ovo" {
		fee = float32((req.Jml + admin) * 0.0318)
	}

	if strings.ToLower(req.Method) == "shopeepay" {
		fee = float32((req.Jml + admin) * 0.04)
	}

	if fee == 0 {
		fee = admin
	}

	d := models.Donasi{
		DonaturID: donaturID,
		ProjectID: ProjectID,
		Method:    strings.ToUpper(req.Method),
		Jml:       req.Jml,
		Fee:       fee,
	}

	if IsOpenFund(ProjectID) != nil {
		return nil, fmt.Errorf("waktu pendanaan proyek id %d sudah ditutup atau belum divalidasi oleh admin", ProjectID)
	}

	return ds.Repo.Create(d)
}

func (ds *DonasiService) CreateInvoice(tr *models.Donasi, email string) (*invoice.Invoice, error) {

	if tr == nil {
		return nil, errors.New("tidak ada input yang dikenali")
	}
	createInvoice := *invoice.NewCreateInvoiceRequest(tr.ID, float64(tr.Jml+tr.Fee))
	xndClient := xendit.NewClient(os.Getenv("XENDIT_SKEY"))
	resp, _, err := xndClient.InvoiceApi.CreateInvoice(context.Background()).CreateInvoiceRequest(createInvoice).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when call invoiceapi.createinvoice %s", err.Error())
	}

	return resp, nil
}

func (ds *DonasiService) GetTransaction(id string, donaturID uint) (interface{}, error) {
	tr, err := ds.Repo.FindByUserID(id, donaturID)
	if err != nil {
		return nil, err
	}
	httpreq, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", getInvoiceUrl, tr.ID), nil)
	if err != nil {
		return nil, err
	}

	httpreq.Header.Set("for-user-id", fmt.Sprintf("%d", tr.DonaturID)) // ex:"1"
	key := os.Getenv("XENDIT_SKEY") + ":"

	encodedString := base64.StdEncoding.EncodeToString([]byte(key))
	httpreq.Header.Set("Authorization", "Basic "+encodedString) // Basic yfhecmhgk98cmyfgsdtvwxe

	client := &http.Client{}
	resp, err := client.Do(httpreq)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("response kode = %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var response interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return response, nil
}

func (ds *DonasiService) UpdateStatus(id string, sts string) error {
	return ds.Repo.UpdateStatus(id, sts)
}

func (ds *DonasiService) ConfirmPayment(id string) error {

	d, err := ds.Repo.ConfirmPayment(id)
	if err != nil {
		return err
	}

	ps := NewProjectService()
	if err = ps.Repo.TambahSaldo(d.ProjectID, d.Jml); err != nil {
		return err
	}
	return nil
}

func (ds *DonasiService) GetAllHistory(projectID uint) ([]models.Donasi, error) {
	return ds.Repo.GetHistory(projectID)
}

func (ds *DonasiService) Contributors(projectID uint, limit uint) ([]models.Donasi, error) {
	return ds.Repo.Contributors(projectID, limit)
}

func (ds *DonasiService) MyContribution(donaturID uint, limit uint) ([]models.Donasi, error) {
	return ds.Repo.MyContribution(donaturID, limit)
}

func (ds *DonasiService) MyHistory(donaturID uint) ([]models.Donasi, error) {
	return ds.Repo.MyHistory(donaturID)
}
