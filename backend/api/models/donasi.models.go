package models

import (
	"errors"
	"time"

	tz "github.com/bagasa11/banturiset/timezone"
	"gorm.io/gorm"
)

type Donasi struct {
	gorm.Model
	ID             string  `gorm:"primaryKey; not null"`
	Status         string  `gorm:"not null"`
	Jml            float32 `gorm:"not null"`
	Fee            float32 `gorm:"not null"`
	Method         string  `gorm:"not null; default:OVO"`
	TrasactionTime time.Time

	ProjectID uint
	Project   Project
	UserID    uint
	User      User
}

func (d *Donasi) BeforeCreate(tx *gorm.DB) error {

	if d.Jml < float32(20000) {
		return errors.New("sumbangan minimum 20k")
	}
	if d.Jml+d.Fee > float32(2000000) {
		return errors.New("total sumbangan maksimum (jml + fee): 2jt")
	}

	tx.Statement.SetColumn("status", "PENDING")
	tx.Statement.SetColumn("TrasactionTime", tz.GetTime(time.Now()))
	return nil
}

func (d *Donasi) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("TrasactionTime", tz.GetTime(time.Now()))
	return nil
}

func (d *Donasi) BeforeDelete(tx *gorm.DB) error {

	return errors.New("tidak boleh menghapus transaksi")
}
