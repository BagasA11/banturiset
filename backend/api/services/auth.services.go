package services

import (
	"errors"
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

	var roleID uint
	user, err := as.Repo.EmailLogin(req.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("gagal mengambil data user")
	}
	// compare input password and hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", nil
	}
	// select user detail by role
	// if admin := user->admin{}
	user, roleID, err = as.selectByRole(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	if roleID == 0 {
		return "", errors.New("terjadi kesalahan dalam fetch data user")
	}

	token, err := helpers.GenerateToken(user.ID, user.Email, user.Role, roleID)
	if err != nil {
		return "", err
	}
	return token, nil

}

func (as *AuthService) selectByRole(userID uint, r string) (*models.User, uint, error) {

	if strings.ToLower(r) == models.Admin {
		return as.Repo.AdminProfile(userID)
	}
	if strings.ToLower(r) == models.Researcher {
		return as.Repo.PenelitiProfile(userID)
	}
	if strings.ToLower(r) == models.Sponsor {
		return as.Repo.DonaturProfile(userID)
	}
	return nil, 0, errors.New("role invalid")
}
