package models

import (
	"gorm.io/gorm"
)

// type Kategori struct {
// 	gorm.Model
// 	ID           uint   `gorm:"primaryKey"`
// 	Judul        string `gorm:"size:20; unique"`
// 	Desk         string
// 	GroupUrl     string
// 	CreatedAt    time.Time
// 	ValidUntil   time.Time
// 	PenyuntingID uint

// 	Penyunting Penyunting `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; embedded"`
// 	Project    []Project
// }

type Kategori struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Title        string `gorm:"size:20; unique"`
	Desc         string
	Link         string `gorm:"size:120"`
	PenyuntingID uint
	Penyunting
}
