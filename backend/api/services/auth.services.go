package services

import (
	"fmt"
	"slices"
	"strings"

	"github.com/bagasa11/banturiset/api/dto"

	"errors"

	"github.com/bagasa11/banturiset/api/repository"
	"github.com/bagasa11/banturiset/helpers"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepo
}

func NewAuthService() *AuthService {
	return &AuthService{
		repo: repository.NewUserRepo(),
	}
}

func (as *AuthService) Login(req dto.Login) (string, error) {

	u, err := as.repo.WhereVerified(req.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return "", nil
	}
	if !u.IsVerfied {
		return "", errors.New("akun anda belum diverifikasi oleh admin")
	}

	if u.IsbBlock {
		return "", errors.New("akun ini telah diblockir")
	}
	r := ""

	if slices.Contains([]string{"penyunting", "reviewer", "admin", "penelaah"}, strings.ToLower(u.Role)) {
		r = "penyunting"
	}
	if slices.Contains([]string{"peneliti", "scientist", "researcher"}, strings.ToLower(u.Role)) {
		r = "peneliti"
	} else {
		r = "donatur"
	}
	fmt.Println("role: ", r)
	token, err := helpers.GenerateToken(u.ID, u.Email, r)
	if err != nil {
		return "", err
	}
	return token, nil

}
