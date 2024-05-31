package models

import (
	"time"

	"gorm.io/gorm"
)

type Kategori struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Judul        string `gorm:"size:20; unique"`
	Desk         string
	GroupUrl     string
	CreatedAt    time.Time
	ValidUntil   time.Time
	PenyuntingID uint

	Penyunting Penyunting `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Project    []Project
}
