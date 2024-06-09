package models

type Progress struct {
	ID          uint   `gorm:"primaryKey"`
	FileUrl     string `gorm:"not null"`
	Desc        string
	Tahap       uint8 `gorm:"not null; default:1"`
	Status      int8  `gorm:"not null; default:0"`
	PesanRevisi *string

	ProjectID uint
	Project   Project
}
