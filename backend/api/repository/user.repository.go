package repository

import (
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
