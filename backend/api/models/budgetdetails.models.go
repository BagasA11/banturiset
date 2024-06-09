package models

import (
	"gorm.io/gorm"
)

type BudgetDetails struct {
	ID        uint `gorm:"primaryKey"`
	Deskripsi string
	Cost      float32 `gorm:"not null; default:0"`

	ProjectID uint
	Project   Project
}

func (bd *BudgetDetails) BeforeCreate(tx *gorm.DB) error {

	return nil
}
