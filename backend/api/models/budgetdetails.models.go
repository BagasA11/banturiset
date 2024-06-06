package models

import (
	"errors"

	"gorm.io/gorm"
)

type BudgetDetails struct {
	ID        uint `gorm:"primaryKey"`
	Deskripsi string
	Tahap     uint8   `gorm:"not null; default:1"`
	Cost      float32 `gorm:"not null; default:0"`
	ProjectID uint
	Project   Project
}

func (bd *BudgetDetails) BeforeCreate(tx *gorm.DB) error {
	if bd.Tahap > uint8(bd.Project.Milestone) {
		return errors.New("tahap pendanaan tidak boleh lebih besar drpd milestone proyek")
	}
	return nil
}
