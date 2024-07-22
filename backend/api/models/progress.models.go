package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Progress struct {
	ID          uint   `gorm:"primaryKey"`
	FileUrl     string `gorm:"not null"`
	Desc        string
	Tahap       uint8 `gorm:"not null; default:1"`
	Validasi    uint8 `gorm:"not null"`
	CreatedAt   time.Time
	ValidatedAt *time.Time

	ValidatorID *uint
	Penyunting  Penyunting `gorm:"foreignKey:ValidatorID"`
	ProjectID   uint
	Project     Project
}

func (p *Progress) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("Validasi", Diajuakan)
	tx.Statement.SetColumn("CreatedAt", time.Now())
	return nil
}

func (p *Progress) BeforeDelete(tx *gorm.DB) error {
	if p.Validasi == Verifikasi {
		return errors.New("tidak boleh menghapus laporan yang sudah divalidasi")
	}
	return nil
}
