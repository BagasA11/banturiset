package services

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
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

	user := models.User{
		FName:         req.FName,
		Email:         req.Email,
		Password:      req.Password,
		Phone:         req.Phone,
		Role:          req.Role,
		Institute:     req.Institute,
		InstituteAddr: req.InstAddr,
		PostCode:      req.PostCode,
	}

	return us.User.Create(user)
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
		return errors.New("role ditolak")
	}
	return us.User.CheckID(id, r)
}

func (us *UserService) Verifikasi(userID uint) (string, error) {
	return us.Penyunting.Verifikasi(userID)
}

func (us *UserService) NotVerified(page uint) ([]models.User, error) {
	return us.Penyunting.NotVerified(page)
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

func (us *UserService) GetProfile(id uint, role string) (*models.User, error) {
	u := new(models.User)
	var err error

	if slices.Contains([]string{"donatur", "investor", "dermawan", "sponsor"}, strings.ToLower(role)) {
		u, err = us.User.DonaturProfile(id)
	}
	if strings.ToLower(role) == "peneliti" {
		u, err = us.User.PenelitiProfile(id)
	}
	if slices.Contains([]string{"admin", "penyunting"}, strings.ToLower(role)) {
		u, err = us.User.AdminProfile(id)
	}

	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("gagal mengambil data user")
	}
	return u, nil
}
