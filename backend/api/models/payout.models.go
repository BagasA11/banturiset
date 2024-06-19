package models

type Payout struct {
	ID         string `gorm:"primaryKey"`
	Tahap      uint8  `gorm:"not null"`
	ProjectID  uint
	PenelitiID uint
	Status     string `gorm:"not null"`
}
