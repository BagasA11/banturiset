package models

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	FName         string `gorm:"not null"`
	Email         string `gorm:"type:string; size:20; not null; unique; <-:create"`
	Password      string `gorm:"not null"`
	Phone         string `gorm:"not null; unique"`
	Role          string `gorm:"index; not null; <-:create"`
	Institute     string `gorm:"not null"`
	InstituteAddr string `gorm:"not null"`
	PostCode      string `gorm:"not null; size:7"`
	Bank          *string
	NoRek         *string
	ProfileUrl    *string
	IsVerfied     bool       `gorm:"not null; default:false"`
	IsBlock       bool       `gorm:"not null; default:false"`
	Peneliti      Peneliti   `gorm:"foreignKey:UserID"`
	Penyunting    Penyunting `gorm:"foreignKey:UserID"`
	Donatur       Donatur    `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("Password", hash)
	if strings.ToLower(u.Role) == "admin" {
		return errors.New("tidak dapat mendaftarkan diri sebagai admin")
	}
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
	return errors.New("tidak boleh menghapus data user")
}

type PaymentInfo struct {
	Bank          string
	AccountNumber string
	Phone         string
}
