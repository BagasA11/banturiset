package repository

import (
	"fmt"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type PenyuntingRepo struct {
	DB *gorm.DB
}

func NewPenyuntingRepo() *PenyuntingRepo {
	return &PenyuntingRepo{
		DB: config.GetDB(),
	}
}

func (ur *PenyuntingRepo) NotVerified(page uint) ([]models.User, error) {
	var users []models.User
	var s = []string{"id", "f_name", "email", "role", "phone", "institute", "post_code"}
	var to = page * 20
	err := ur.DB.Where("is_verfied = ?", false).Where("id BETWEEN ? AND ?", to-19, to).Select(s).Find(&users).Error
	return users, err
}

func (ur *PenyuntingRepo) Verifikasi(id uint) (string, error) {
	var u models.User
	if err := ur.DB.Where("is_verfied = ? AND is_block = ?", false, false).First(&u, id).Error; err != nil {
		return "", fmt.Errorf("gagal melakukan verifikasi. error: %s", err.Error())
	}

	u.IsVerfied = true
	ur.DB.Save(&u)
	return u.Email, nil
}

func (ur *PenyuntingRepo) Blokir(id uint) error {

	tx := ur.DB.Begin()
	if err := tx.Model(&models.User{}).Where("id = ?", id).Update("IsBlock", true).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
