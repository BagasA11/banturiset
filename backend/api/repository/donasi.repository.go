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

func (dr *DonasiRepo) ConfirmPayment(id string) (models.Donasi, error) {
	var d models.Donasi
	if err := dr.DB.Preload("Project").Where("id = ?", id).First(&d).Error; err != nil {

		fmt.Println("error dr->confirmPayment(): ", err.Error())
		if err == gorm.ErrRecordNotFound {
			return d, gorm.ErrRecordNotFound
		}
		return d, errors.New("error when confirm payment status")
	}
	d.Status = "PAID"
	err := dr.DB.Save(&d).Error
	return d, err
}

func (dr *DonasiRepo) UpdateStatus(id string, status string) error {
	fmt.Printf("\nid: %s \t status: %s\n", id, status)
	tx := dr.DB.Begin()
	if err := tx.Model(&models.Donasi{}).Where("id = ?", id).Update("status", status).Error; err != nil {
		tx.Rollback()
		fmt.Println("error: ", err.Error())
		return err
	}
	tx.Commit()
	return nil
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
	err := dr.DB.
		Where("status = ? AND project_id = ?", "PAID", projectID).
		Limit(int(limit)).Order("jml desc").
		Find(&d).Error
	return d, err
}

func (dr *DonasiRepo) MyContribution(donaturID uint, limit uint) ([]models.Donasi, error) {
	var d []models.Donasi
	if err := dr.DB.
		Where("status = ? AND donatur_id = ?", "PAID", donaturID).
		Preload("Project").
		Limit(int(limit)).
		Find(&d).Error; err != nil {

		fmt.Println("error dr->myContribution(): ", err.Error())
		return d, errors.New("gagal mengambil data")
	}
	return d, nil
}

func (dr *DonasiRepo) MyHistory(donaturID uint) ([]models.Donasi, error) {
	var d []models.Donasi
	if err := dr.DB.Order("updated_at DESC").Where("donatur_id = ?", donaturID).Preload("Project").
		Find(&d).Error; err != nil {
		fmt.Println("error dr->myContribution(): ", err.Error())
		return d, errors.New("gagal mengambil data")
	}
	return d, nil
}
