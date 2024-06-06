package repository

import (
	"errors"
	"fmt"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		DB: config.GetDB(),
	}
}

func (ur *UserRepo) Create(u models.User) (uint, error) {
	tx := ur.DB.Begin()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return u.ID, nil
}

func (ur *UserRepo) FindID(id uint) (models.User, error) {
	var user models.User
	err := ur.DB.Where("id = ?", false).First(&user).Error
	return user, err
}

func (ur *UserRepo) CheckID(id uint, role string) error {
	return ur.DB.Select("ID").Where("role = ?", role).First(&models.User{}, id).Error
}

func (ur *UserRepo) SetAvatar(id uint, url string) error {
	tx := ur.DB.Begin()
	if err := tx.Model(&models.User{}).Where("id = ?", id).Update("profile_url", url).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ur *UserRepo) WhereVerified(email string) (models.User, error) {
	var u models.User
	if err := ur.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur *UserRepo) AdminLogin(email string) (models.User, error) {
	fmt.Println("\nemail: ", email)
	var admin models.User
	if err := ur.DB.Where("email = ? AND is_verfied = ? AND is_block = ?", email, true, false).Joins("Penyunting").First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}
func (ur *UserRepo) PenelitLogin(email string) (models.User, error) {
	var peneliti models.User
	if err := ur.DB.Where("email = ? AND is_verfied = ? AND is_block = ?", email, true, false).Joins("Peneliti").First(&peneliti).Error; err != nil {
		return peneliti, err
	}
	return peneliti, nil
}

func (ur *UserRepo) DonaturLogin(email string) (models.User, error) {
	var donatur models.User
	if err := ur.DB.Where("email = ? AND is_verfied = ? AND is_block = ?", email, true, false).Joins("Donatur").First(&donatur).Error; err != nil {
		return donatur, err
	}
	return donatur, nil

}

func (pr *UserRepo) Update(user *models.User) error {
	if user == nil {
		return errors.New("informasi update kosong")
	}
	tx := pr.DB.Begin()
	if err := tx.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
