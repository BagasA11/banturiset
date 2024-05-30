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
