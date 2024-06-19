package services

import (
	"errors"
	"slices"
	"strings"

	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
)

type DonasiService struct {
	Repo *repository.DonasiRepo
}

const paymenturl = "https://api.xendit.co/v2/invoices"
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
		return nil, fmt.Errorf("waktu pendanaan proyek id %d sudah ditutup", ProjectID)
	}

	return ds.Repo.Create(d)
}

func (ds *DonasiService) CreateInvoice(tr *models.Donasi, email string) (*dto.InvoicePage, error) {

	if tr == nil {
		return nil, errors.New("tidak ada input yang dikenali")
	}

	// create body payload
	fee := map[string]interface{}{
		"type":  "ADMIN",
		"value": tr.Fee,
	}

	data := map[string]interface{}{
		"external_id": tr.ID,
		"for-user-id": fmt.Sprintf("%d", tr.DonaturID),
		"amount":      tr.Jml,
		"currency":    "IDR",
		"customer": map[string]interface{}{
			"email": email,
		},
		"customer_notification_preference": map[string]interface{}{
			"invoice_paid":    "email",
			"invoice_created": "email",
		},
		"payment_methods": tr.Method,
		"fees":            []interface{}{fee},
	}

	// encode payload to json
	jsondata, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", paymenturl, bytes.NewBuffer(jsondata))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	key := os.Getenv("XENDIT_SKEY") + ":"
	encodedString := base64.StdEncoding.EncodeToString([]byte(key))
	httpReq.Header.Set("Authorization", encodedString)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err

	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("response status = %d", resp.StatusCode)
	}

	var createTransactionResponse *dto.InvoicePage
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&createTransactionResponse); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return createTransactionResponse, nil
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
	key := os.Getenv("XENDIT_SKEY") + ":"                              // ex: xxxxx:
	fmt.Println("skey: ", key)
	encodedString := base64.StdEncoding.EncodeToString([]byte(key))
	httpreq.Header.Set("Authorization", encodedString)

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

func (ds *DonasiService) Notifikasi(req dto.NotifInvoice) error {
	d, err := ds.Repo.FindID(req.ExternalID)
	if err != nil {
		return err
	}

	if err := ds.Repo.UpdateStatus(req.ExternalID, req.Status); err != nil {
		return err
	}

	if strings.ToLower(req.Status) != "paid" {
		return nil
	}
	bs := NewProjectService()
	return bs.Repo.TambahSaldo(d.ProjectID, d.Jml)
}
