package models

import (
	"errors"
	"time"

	tz "github.com/bagasa11/banturiset/timezone"
	"gorm.io/gorm"
)

type Report struct {
	ID      uint   `gorm:"primaryKey"`
	FileUrl string `gorm:"not null"`
	// Milestone uint8
	Desc string
	// Versi       uint8 `gorm:"not null; default:1"`
	Validasi    uint8 `gorm:"not null"`
	CreatedAt   time.Time
	ValidatedAt *time.Time

	ValidatorID *uint
	Penyunting  Penyunting `gorm:"foreignKey:ValidatorID"`
	ProjectID   uint
	Project     Project
}

func (p *Report) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("Validasi", Diajuakan)
	tx.Statement.SetColumn("CreatedAt", tz.GetTime(time.Now()))
	return nil
}

func (p *Report) BeforeDelete(tx *gorm.DB) error {
	if p.Validasi == Verifikasi {
		return errors.New("tidak boleh menghapus laporan yang sudah divalidasi")
	}
	return nil
}
