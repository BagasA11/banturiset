package repository

import (
	"errors"

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

func (dr *DonasiRepo) UpdateStatus(id string, status string) (models.Donasi, error) {
	tx := dr.DB.Begin()
	// fetch data from db
	var d models.Donasi
	if err := dr.DB.Where("id = ?", id).First(&d).Error; err != nil {
		fmt.Println("error ketika mengambil data: ", err.Error())
		return d, errors.New("error when update status")
	}

	d.Status = status
	if err := tx.Where("id = ?", d.ID).Save(&d).Error; err != nil {
		tx.Rollback()
		fmt.Println("error", err.Error())
		return d, errors.New("error when update status")
	}
	if err := tx.Commit().Error; err != nil {
		fmt.Println("gagal melakukan rollback", err.Error())
		return d, errors.New("error when update status")
	}
	return d, nil
}

func (dr *DonasiRepo) FindByUserID(trID string, donaturID uint) (models.Donasi, error) {
	var d models.Donasi
	err := dr.DB.Where("id = ? AND donatur_id = ?", trID, donaturID).First(&d).Error
	return d, err
}

func (dr *DonasiRepo) FindID(trID string) (models.Donasi, error) {
	var d models.Donasi
	err := dr.DB.First(&d, trID).Error
	return d, err
}

func (dr *DonasiRepo) GetHistory(projectID uint) ([]models.Donasi, error) {
	var d []models.Donasi
	if err := dr.DB.Where("project_id = ?", projectID).Find(&d).Error; err != nil {
		fmt.Println("error: ", err.Error())
		return d, errors.New("gagal mendapatkan histori donasi")
	}
	return d, nil
}

func (dr *DonasiRepo) Contributors(projectID uint, limit uint) ([]models.Donasi, error) {
	var d []models.Donasi
	// select * from donasis WHERE status = 'PAID' AND project_id = $projectID order by
	err := dr.DB.Where("status = ? AND project_id = ?", "PAID", projectID).Limit(int(limit)).Order("jml desc").Find(&d).Error
	return d, err
}
