package models

import (
	"errors"

	"time"

	"gorm.io/gorm"
)

// struktur urutan token
// id, token, expired_date, isblakclist
type TokenList struct {
	ID          uint   `gorm:"primaryKey"`
	Token       string `gorm:"unique; not null"`
	ExpiredDate time.Time
	Blacklisted bool
}

func (tl *TokenList) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("Blacklisted", false)
	return nil
}

func (tl *TokenList) BeforeDelete(tx *gorm.DB) error {
	return errors.New("token tidak dapat dihapus")
}
