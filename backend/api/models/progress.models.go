package models

import (
	"errors"

	"gorm.io/gorm"
)

type Progress struct {
	ID       uint   `gorm:"primaryKey"`
	FileUrl  string `gorm:"not null"`
	Desc     string
	Tahap    uint8 `gorm:"not null; default:1"`
	Validasi uint8 `gorm:"not null"`

	ProjectID uint
	Project   Project
}

func (p *Progress) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("Validasi", Diajuakan)
	return nil
}

func (p *Progress) BeforeDelete(tx *gorm.DB) error {
	if p.Validasi == Verifikasi {
		return errors.New("tidak boleh menghapus laporan yang sudah divalidasi")
	}
	return nil
}
