package models

type Donatur struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
}
