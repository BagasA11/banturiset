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

	// mencari proyek dengan id = $id, waktu pendanaan masih dibuka, dan diverifikasi dan diblokir = false
	if tx.Where("fund_until >= ? AND status >= ? AND is_block = ", time.Now(), models.Verifikasi, false).First(&models.Project{}, d.ProjectID) != nil {
		return nil, errors.New("waktu pendanaan proyek ini sudah ditutup")
	}

	if err := tx.Create(&d).Error; err != nil {
		tx.Rollback()
		return nil, err
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

func (dr *DonasiRepo) Contributors(projectID uint) ([]models.Donasi, error) {
	var d []models.Donasi
	// select * from donasis WHERE status = 'PAID' AND project_id = $projectID order by
	err := dr.DB.Where("status = ? AND project_id = ?", "PAID", projectID).Order("jml desc").Find(&d).Error
	return d, err
}
