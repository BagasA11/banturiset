package repository

import (
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type DonaturRepo struct {
	DB *gorm.DB
}

func NewDonaturRepo() *DonaturRepo {
	return &DonaturRepo{
		DB: config.GetDB(),
	}
}

func (dr *DonaturRepo) Create(u models.Donatur) error {
	tx := dr.DB.Begin()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
