package models

import "gorm.io/gorm"

type Payout struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	ProjectID  uint
	PenelitiID uint
	Message    string
	Status     bool `gorm:"not null; default:false"`
}
