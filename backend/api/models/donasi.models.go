package models

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Donasi struct {
	gorm.Model
	ID     string  `gorm:"primaryKey; not null"`
	Status string  `gorm:"not null"`
	Jml    float32 `gorm:"not null"`
	Fee    float32 `gorm:"not null"`
	Method string  `gorm:"not null; default:OVO"`

	ProjectID uint
	Project   Project
	DonaturID uint
	Donatur   Donatur
}

func (d *Donasi) BeforeCreate(tx *gorm.DB) error {

	if d.Jml < float32(20000) {
		return errors.New("sumbangan minimum 20k")
	}
	tx.Statement.SetColumn("status", "PENDING")
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (d *Donasi) BeforeUpdate(tx *gorm.DB) error {
	if strings.ToLower(d.Status) == "paid" {
		return errors.New("tidak dapat mengubah data transaksi yang sudah berhasil dibayar")
	}
	return nil
}

func (d *Donasi) BeforeDelete(tx *gorm.DB) error {

	return errors.New("tidak boleh menghapus transaksi")
}
