package models

import (
	"gorm.io/gorm"
)

type Penyunting struct {
	gorm.Model
	NIP    string `gorm:"type:string; primaryKey"`
	UserID uint
}
