package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type ReportProgressRepo struct {
	DB *gorm.DB
}

func NewProgressRepo() *ReportProgressRepo {
	return &ReportProgressRepo{
		DB: config.GetDB(),
	}
}

func (rp *ReportProgressRepo) CreateReport(p models.Progress) error {
	tx := rp.DB.Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
		return errors.New("gagal menambah data laporan ke database")
	}
	tx.Commit()
	return nil
}

func (rp *ReportProgressRepo) IsRedundant(projectID uint, tahap uint8) error {
	var p []models.Progress
	if err := rp.DB.Where("project_id = ? AND tahap = ?", projectID, tahap).Find(&p).Error; err != nil {
		fmt.Printf("error progress->isRedundant(): %s", err.Error())
		return errors.New("gagal mengambil data")
	}
	if len(p) == 0 {
		return nil
	}
	return fmt.Errorf("tahap ke-%d sudah didefinisikan", p[0].Tahap)
}
