package services

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
	ef "github.com/bagasa11/banturiset/errorf"
)

type UserService struct {
	User       *repository.UserRepo
	Peneliti   *repository.PenelitiRepo
	Donatur    *repository.DonaturRepo
	Penyunting *repository.PenyuntingRepo
}

func NewUserService() *UserService {
	return &UserService{
		User:       repository.NewUserRepo(),
		Peneliti:   repository.NewPenelitiRepo(),
		Donatur:    repository.NewDonaturRepo(),
		Penyunting: repository.NewPenyuntingRepo(),
	}
}

func (us *UserService) UserRegister(req dto.UserRegister) (uint, error) {
	if strings.Contains(strings.ToLower(req.Role), "admin") {
		return 0, errors.New("admin tidak boleh melakukan registrasi")
	}

	if !slices.Contains([]string{models.Sponsor, models.Researcher}, strings.ToLower(req.Role)) {
		return 0, fmt.Errorf("tipe user invalid: %v", []string{models.Sponsor, models.Researcher})
	}

	user := autoValidateUser(req, req.Role)

	return us.User.Create(user)
}

// @param dto.UserRegister
// @param string
// if role is donatur, then user verified property is true. So User can be login immadiately after register proccess
// @return models.User
func autoValidateUser(req dto.UserRegister, role string) models.User {

	if strings.ToLower(role) == models.Sponsor {
		return models.User{
			FName:         req.FName,
			Email:         req.Email,
			Password:      req.Password,
			Phone:         req.Phone,
			Role:          req.Role,
			Institute:     req.Institute,
			InstituteAddr: req.InstAddr,
			PostCode:      req.PostCode,
			IsVerfied:     true,
		}
	}

	return models.User{
		FName:         req.FName,
		Email:         req.Email,
		Password:      req.Password,
		Phone:         req.Phone,
		Role:          req.Role,
		Institute:     req.Institute,
		InstituteAddr: req.InstAddr,
		PostCode:      req.PostCode,
		IsVerfied:     false,
	}
}

func (us *UserService) CreateDonatur(userID uint) error {
	d := models.Donatur{
		UserID: userID,
	}

	rd, err := us.Donatur.IsRedundant(d.UserID)
	if err != nil {
		return err
	}
	if !rd {
		return fmt.Errorf("peneliti dengan userID: %d sudah terdaftar", d.UserID)
	}

	if err := us.User.IsPeneliti(userID); err != nil {
		return err
	}

	return us.Donatur.Create(d)
}

func (us *UserService) CreatePeneliti(userID uint, req dto.PenelitiRegister) error {
	p := models.Peneliti{
		NIP:    req.NIP,
		UserID: userID,
	}
	// redundant check
	// memeriksa peneliti yang memiliki userID sama
	rd, err := us.Peneliti.IsRedundant(p.UserID)
	if err != nil {
		return err
	}
	if !rd {
		return fmt.Errorf("peneliti dengan userID: %d sudah terdaftar", p.UserID)
	}

	if err := us.User.IsDonatur(userID); err != nil {
		return err
	}

	return us.Peneliti.Create(p)
}

func (us *UserService) CheckID(id uint, role string) error {
	r := strings.ToLower(role)
	if !slices.Contains([]string{"donatur", "peneliti", "researcher", "saintist"}, strings.ToLower(r)) {
		return ef.ErrInvalidRole
	}
	return us.User.CheckID(id, r)
}

func (us *UserService) Verifikasi(userID uint) (string, error) {
	return us.Penyunting.Verifikasi(userID)
}

func (us *UserService) NotVerified(page uint) ([]models.User, error) {
	var last = page * 20
	var begin = last - 19
	fmt.Printf("%d - %d\n", begin, last)
	return us.Penyunting.NotVerified(begin, last)
}

func (us *UserService) CompletePayentInfo(id uint, req dto.PaymentInfos) error {

	if !slices.Contains([]string{"bca", "bsi", "mandiri", "bri", "bni", "bjb"}, strings.ToLower(req.Bank)) {
		return fmt.Errorf("hanya menerima provider bank %v",
			fmt.Sprintf("hanya menerima provider bank %v", []string{"bca", "bsi", "mandiri", "bri", "bni", "bjb"}))
	}
	bank := strings.ToUpper(req.Bank)
	u := models.User{
		ID:    id,
		Bank:  &bank,
		NoRek: &req.NoRek,
	}

	return us.User.Update(&u)
}

func (us *UserService) GetProfile(id uint, role string) (*models.User, uint, error) {
	return us.selectByRole(id, role)
}

func GetPaymentInfo(userID uint) (models.PaymentInfo, error) {
	us := NewUserService()
	return us.User.GetPayment(userID)
}

func (us *UserService) selectByRole(userID uint, r string) (*models.User, uint, error) {

	if strings.ToLower(r) == models.Admin {
		return us.User.AdminProfile(userID)
	}
	if strings.ToLower(r) == models.Researcher {
		return us.User.PenelitiProfile(userID)
	}
	if strings.ToLower(r) == models.Sponsor {
		return us.User.DonaturProfile(userID)
	}
	return nil, 0, errors.New("role invalid")
}

func (us *UserService) ReviewProfile(id uint) (models.User, error) {
	u, err := us.User.ReviewProfile(id)
	if err != nil {
		return u, err
	}
	// d->uID = p->uID = pny->uID = 0
	if u.Donatur.UserID == u.Peneliti.UserID && u.Peneliti.UserID == u.Penyunting.UserID && u.Penyunting.UserID == 0 {
		return u, errors.New("data Detail User belum diinput")
	}
	return u, nil
}

// 1 0 0 -> 1
// 0 1 0 -> 1

// 0 0 0 -> 0
