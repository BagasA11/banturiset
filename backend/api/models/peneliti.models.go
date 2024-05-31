package models

import "gorm.io/gorm"

type Peneliti struct {
	gorm.Model
	NIP    string `gorm:"type:string; primaryKey; <-:create"`
	UserID uint

	Project []Project
}
