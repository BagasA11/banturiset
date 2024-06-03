package repository

import (
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type PenelitiRepo struct {
	DB *gorm.DB
}

func NewPenelitiRepo() *PenelitiRepo {
	return &PenelitiRepo{
		DB: config.GetDB(),
	}
}

func (pr *PenelitiRepo) Create(p models.Peneliti) error {
	tx := pr.DB.Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *PenelitiRepo) IsRedundant(userID uint) (bool, error) {
	var peneliti []models.Peneliti
	if err := pr.DB.Where("user_id = ?", userID).Find(&peneliti).Error; err != nil {
		return false, err
	}
	// sudah daftar, berarti sudah ada userID
	// len(peneliti) > 0
	if len(peneliti) > 0 {
		return false, nil
	}
	return true, nil
}
