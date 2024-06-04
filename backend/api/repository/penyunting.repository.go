package repository

import (
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
	err := ur.DB.Where("is_verfied = ?", false).Where("id BETWEEN ? AND ?", page, page+10).Find(&users).Error
	return users, err
}

func (ur *PenyuntingRepo) Verifikasi(id uint) (string, error) {
	var u models.User
	tx := ur.DB.Begin()
	if err := tx.Model(&u).Where("id = ?", id).Update("is_verfied", true).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
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
