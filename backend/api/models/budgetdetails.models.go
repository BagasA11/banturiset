package models

import (
	"gorm.io/gorm"
)

type BudgetDetails struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	Cost      float32 `gorm:"not null; default:'0' "`
	Desc      string  `gorm:"not null; size:265 "`
	ProjectID uint
	Project   Project `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
