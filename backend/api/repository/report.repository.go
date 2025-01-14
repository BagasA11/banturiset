package repository

import (
	"errors"
	"log"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type ReportRepo struct {
	DB *gorm.DB
}

func NewReportRepo() *ReportRepo {
	return &ReportRepo{
		DB: config.GetDB(),
	}
}

func (rp *ReportRepo) CreateReport(p models.Report) error {
	tx := rp.DB.Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
		return errors.New("gagal menambah data laporan ke database")
	}
	tx.Commit()
	return nil
}

// memeriksa redundansi pada Tahap laporan
// Jika Tahap ditemukan, atau error saat mengambil data maka akan me return error
// jika Record Not Found, maka akan me-return nil
// func (rp *ReportRepo) IsRedundant(projectID uint, tahap uint8) error {
// 	var p models.Report
// 	if err := rp.DB.Where("project_id = ? AND tahap = ?", projectID, tahap).First(&p).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil
// 		}
// 		return errors.New("gagal mengambil data")
// 	}

// 	return fmt.Errorf("tahap ke-%d sudah didefinisikan", p.Tahap)
// }
