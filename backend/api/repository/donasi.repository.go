package repository

import (
	"errors"
	"time"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"

	"fmt"

	"github.com/bagasa11/banturiset/helpers"
)

type DonasiRepo struct {
	DB *gorm.DB
}

func NewDonasiRepo() *DonasiRepo {
	return &DonasiRepo{
		DB: config.GetDB(),
	}
}

func (dr *DonasiRepo) Create(d models.Donasi) (*models.Donasi, error) {
	d.ID = fmt.Sprintf("invoice-%s", helpers.RandStr(7))
	tx := dr.DB.Begin()
	if err := tx.Create(&d).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if time.Now().After(d.Project.FundUntil) {
		return nil, errors.New("pendanaan sudah ditutup")
	}

	tx.Commit()
	return &d, nil
}

func (dr *DonasiRepo) ConfirmPayment(id uint) (float32, error) {
	d := models.Donasi{}
	tx := dr.DB.Begin()
	if err := tx.First(&d, id).Error; err != nil {
		return float32(0), err
	}

	d.Status = "PAID"
	if err := tx.Save(&d).Error; err != nil {
		tx.Rollback()
		return float32(0), err
	}
	tx.Commit()
	return d.Jml, nil
}

func (dr *DonasiRepo) UpdateStatus(id string, status string) error {
	tx := dr.DB.Begin()
	if err := tx.Model(&models.Donasi{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (dr *DonasiRepo) FindByUserID(trID string, donaturID uint) (models.Donasi, error) {
	var d models.Donasi
	err := dr.DB.Where("donatur_id = ?", donaturID).First(&d, trID).Error
	return d, err
}

func (dr *DonasiRepo) FindID(trID string) (models.Donasi, error) {
	var d models.Donasi
	err := dr.DB.First(&d, trID).Error
	return d, err
}
