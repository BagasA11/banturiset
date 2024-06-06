package services

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
	"github.com/bagasa11/banturiset/helpers"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepo
}

func NewAuthService() *AuthService {
	return &AuthService{
		Repo: repository.NewUserRepo(),
	}
}

func (as *AuthService) Login(req dto.Login) (string, error) {

	var u models.User
	var err error
	var roleID uint

	if req.Role == "" {
		return "", errors.New("parameter kosong")
	}

	if !slices.Contains([]string{"admin", "penyunting", "peneliti", "donatur", "penyumbang", "investor", "sponsor", "dermawan"}, strings.ToLower(req.Role)) {
		return "", errors.New("role invalid, pilih role tersedia:{admin, penyunting, peneliti, donatur}")
	}

	// admin
	if slices.Contains([]string{"admin", "penyunting"}, strings.ToLower(req.Role)) {
		u, err = as.Repo.AdminLogin(req.Email)
		roleID = u.Penyunting.ID
	}
	// penelit
	if strings.ToLower(req.Role) == "peneliti" {
		u, err = as.Repo.PenelitLogin(req.Email)
		roleID = u.Peneliti.ID
	}
	// donatur
	if slices.Contains([]string{"donatur", "investor", "dermawan", "sponsor"}, strings.ToLower(req.Role)) {
		u, err = as.Repo.DonaturLogin(req.Email)
		roleID = u.Donatur.ID
	}

	if err != nil {
		return "", fmt.Errorf("gagal login email:%s", req.Email)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return "", nil
	}

	token, err := helpers.GenerateToken(u.ID, u.Email, u.Role, roleID)
	if err != nil {
		return "", err
	}
	return token, nil

}
