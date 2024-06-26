package models

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	FName         string
	Email         string `gorm:"type:string; size:20; not null; unique"`
	Password      string `gorm:"not null"`
	Phone         string `gorm:"not null"`
	Role          string `gorm:"index not null"`
	Institute     string `gorm:"not null"`
	InstituteAddr string `gorm:"not null"`
	PostCode      string `gorm:"not null"`
	Bank          *string
	ProfileUrl    *string
	IsVerfied     string `gorm:"not null default:'false'"`
	IsbBlock      string `gorm:"not null default:'false'"`
	Penyunting    Penyunting
	Donatur       Donatur
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
