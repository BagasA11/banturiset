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

	return us.Peneliti.Create(p)
}

func (us *UserService) CheckID(id uint, role string) error {
	r := strings.ToLower(role)
	if !slices.Contains([]string{"donatur", "peneliti", "researcher", "saintist"}, strings.ToLower(r)) {
		return errors.New("role ditolak")
	}
	return us.User.CheckID(id, r)
}

func (us *UserService) Verifikasi(userID uint) error {
	return us.Penyunting.Verifikasi(userID)
}

func (us *UserService) NotVerified(page uint) ([]models.User, error) {
	return us.Penyunting.NotVerified(page)
}
