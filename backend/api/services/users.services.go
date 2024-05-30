package services

import (
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
