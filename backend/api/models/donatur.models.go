package models

import "gorm.io/gorm"

type Donatur struct {
	gorm.Model
	ID     uint `gorm:"primaryKey"`
	UserID uint
}
