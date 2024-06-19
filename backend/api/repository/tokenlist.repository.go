package repository

import (
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type TokenRepo struct {
	DB *gorm.DB
}

func NewTokenListRepo() *TokenRepo {
	return &TokenRepo{
		DB: config.GetDB(),
	}
}

func (tr *TokenRepo) Create(tl models.TokenList) error {
	tx := tr.DB.Begin()
	if err := tx.Create(&tl).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (tr *TokenRepo) Get(token string) (models.TokenList, error) {
	tl := models.TokenList{}
	if err := tr.DB.Where("token = ? AND blacklisted = ?", token, false).First(&tl).Error; err != nil {
		return models.TokenList{}, err
	}
	return tl, nil
}

func (tr *TokenRepo) Blacklist(token string) error {
	tx := tr.DB.Begin()
	if err := tx.Where("token = ?", token).Update("blacklisted", true).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
