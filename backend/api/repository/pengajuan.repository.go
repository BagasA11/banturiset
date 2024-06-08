package repository

import (
	"time"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type PengajuanRepository struct {
	DB *gorm.DB
}

func NewPengajuanRepository() *PengajuanRepository {
	return &PengajuanRepository{
		DB: config.GetDB(),
	}
}

func (pr *PengajuanRepository) Create(p models.Pengajuan) error {
	tx := pr.DB.Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *PengajuanRepository) Open() ([]models.Pengajuan, error) {
	var p []models.Pengajuan
	err := pr.DB.Where("closed_at > ?", time.Now()).Select([]string{"id", "closed_at", "link_panduan", "desc", "title", "icon_url"}).Find(&p).Error
	return p, err
}

func (pr *PengajuanRepository) FindID(id uint) (models.Pengajuan, error) {
	var p models.Pengajuan
	err := pr.DB.First(&p, id).Error
	return p, err
}