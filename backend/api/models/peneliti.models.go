package models

type Peneliti struct {
	// gorm.Model
	ID     uint   `gorm:"primaryKey"`
	NIP    string `gorm:"type:string; <-:create; size:20"`
	UserID uint

	Project []Project
	Payout  []Payout
}
