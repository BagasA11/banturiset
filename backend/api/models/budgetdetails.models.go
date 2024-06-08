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
	Percent   uint8   `gorm:"not null; default:0"`

	ProjectID uint
	Project   Project
}

func (bd *BudgetDetails) BeforeCreate(tx *gorm.DB) error {

	if bd.Percent >= 100 {
		return errors.New("percentase budget tidak boleh mendekati 100%")
	}

	return nil
}
