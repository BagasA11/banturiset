package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Pengajuan struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	Title        string `gorm:"size:20; unique"`
	Desc         string
	LinkWa       string  `gorm:"size:120"`
	LinkPanduan  *string `gorm:"size:120"`
	IconUrl      string  `gorm:"size:120"`
	ClosedAt     time.Time
	PenyuntingID uint
	Penyunting   Penyunting
	Project      []Project
}

func (p *Pengajuan) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("CreatedAt", time.Now())
	// if closed - created < 3 months

	if time.Since(p.ClosedAt).Nanoseconds() < time.Hour.Nanoseconds()*24*30*3 {
		return errors.New("tenggat waktu minimal 3 bulan")
	}
	return nil
}
