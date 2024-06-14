package models

import (
	"errors"

	"gorm.io/gorm"
)

type Progress struct {
	ID          uint   `gorm:"primaryKey"`
	FileUrl     string `gorm:"not null"`
	Desc        string
	Tahap       uint8 `gorm:"not null; default:1"`
	Status      int8  `gorm:"not null; default:0"`
	PesanRevisi *string

	ProjectID uint
	Project   Project
}

func (p *Progress) BeforeDelete(tx *gorm.DB) error {
	if p.Status >= Verifikasi {
		return errors.New("tidak boleh menghapus laporan yang sudah divalidasi")
	}
	return nil
}
