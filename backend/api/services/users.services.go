package services

import (
	"errors"
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
		Phone:         req.Password,
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
	return us.Donatur.Create(d)
}

func (us *UserService) CreatePeneliti(userID uint, req dto.PenelitiRegister) error {
	p := models.Peneliti{
		NIP:    req.NIP,
		UserID: userID,
	}
	return us.Peneliti.Create(p)
}
