package models

import "gorm.io/gorm"

type Peneliti struct {
	gorm.Model
	NIP    string `gorm:"type:string; primaryKey"`
	UserID uint
}
