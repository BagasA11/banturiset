package models

import (
	"errors"

	"gorm.io/gorm"
)

type Penyunting struct {
	ID       uint   `gorm:"primaryKey"`
	NIP      string `gorm:"type:string; <-:create; size:20"`
	UserID   uint
	Kategori []Kategori
}

func (p *Penyunting) BeforeCreate(tx *gorm.DB) error {
	return errors.New("admin tidak boleh melakukan registrasi")
}
