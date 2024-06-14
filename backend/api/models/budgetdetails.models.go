package models

import (
	"errors"

	"gorm.io/gorm"
)

type BudgetDetails struct {
	ID        uint `gorm:"primaryKey"`
	Deskripsi string
	Cost      float32 `gorm:"not null; default:0"`

	ProjectID uint
	Project   Project
}

func (bd *BudgetDetails) BeforeUpdate(tx *gorm.DB) error {
	if bd.Project.Status >= Verifikasi {
		return errors.New("tidak bisa mengubah detil ketika proyek telah diverifikasi")
	}
	return nil
}

func (bd *BudgetDetails) BeforeDelete(tx *gorm.DB) error {
	if bd.Project.Status >= Verifikasi {
		return errors.New("proyek sudah disetujui. tidak dapat mengedit tahapan")
	}
	return nil
}
