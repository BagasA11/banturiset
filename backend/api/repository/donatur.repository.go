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

func (dr *DonaturRepo) IsRedundant(userID uint) (bool, error) {
	var donatur []models.Donatur

	if err := dr.DB.Where("user_id = ?", userID).Find(&donatur).Error; err != nil {
		return false, err
	}
	// sudah daftar, berarti sudah ada userID
	// len(peneliti) > 0
	if len(donatur) > 0 {
		return false, nil
	}
	return true, nil
}
