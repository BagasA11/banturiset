package models

// import (
// 	"errors"

// 	"gorm.io/gorm"
// )

// type BudgetDetails struct {
// 	ID        uint    `gorm:"primaryKey"`
// 	Cost      float32 `gorm:"not null;"`
// 	Desc      string  `gorm:"not null; size:265 "`
// 	ProjectID uint
// 	Project   Project `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
// }

// func (bd *BudgetDetails) BeforeCreate(tx *gorm.DB) error {
// 	if int32(bd.Cost) < 0 {
// 		return errors.New("budget harus > 0")
// 	}
// 	return nil
// }
