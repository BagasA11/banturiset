package repository

import (
	"errors"
	"time"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	tz "github.com/bagasa11/banturiset/timezone"
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
	err := pr.DB.Where("closed_at > ?", tz.GetTime(time.Now())).Select([]string{"id", "closed_at", "link_panduan", "desc", "title", "icon_url"}).Find(&p).Error
	return p, err
}

func (pr *PengajuanRepository) IsOpen(id uint) error {
	if err := pr.DB.Where("closed_at >= ?", tz.GetTime(time.Now())).First(&models.Pengajuan{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return nil
}

func (pr *PengajuanRepository) FindID(id uint) (models.Pengajuan, error) {
	var p models.Pengajuan
	err := pr.DB.First(&p, id).Error
	return p, err
}
