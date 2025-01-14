package models

type Peneliti struct {
	// gorm.Model
	ID                 uint   `gorm:"primaryKey" json:"id"`
	NIP                string `gorm:"type:string; <-:create; size:20" json:"nip"`
	UserID             uint   `json:"user_id"`
	Institute          string `json:"institute" gorm:"type:string; not null"`
	Address            string `json:"address" gorm:"type:string; not null"`
	Phone              string `gorm:"not null; unique"`
	EtherWalletAddress *string
	Project            []Project
	Payout             []Payout
}
