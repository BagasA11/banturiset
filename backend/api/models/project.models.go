package models

import (
	"errors"

	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:50; not null"`
	Desc          string
	ProposalUrl   *string   `gorm:"size:170"`
	KlirensUrl    *string   `gorm:"size:170"`
	FundUntil     time.Time `gorm:"not null"`
	DeadLine      time.Time `gorm:"not null"`
	Milestone     int8      `gorm:"not null; default:1"`
	TktLevel      int8      `gorm:"not null; default:1"`
	Cost          float32   `gorm:"not null;"`
	Status        int8      `gorm:"not null; size:10; default:0 "`
	PesanRevisi   *string
	Fraud         bool    `gorm:"not null; default:false"`
	CollectedFund float32 `gorm:"not null"`

	PengajuanID uint      `gorm:"not null"`
	Pengajuan   Pengajuan `gorm:"foreignKey:PengajuanID"`
	PenelitiID  uint      `gorm:"not null"`
	Peneliti    Peneliti  `gorm:"foreignKey:PenelitiID"`

	Donasi        []Donasi
	BudgetDetails []BudgetDetails
	Tahapan       []Tahapan
	Progress      []Progress
}

func (p *Project) BeforeCreate(tx *gorm.DB) error {

	if p.Cost < float32(0) {
		return errors.New("biaya harus > 0")
	}
	tx.Statement.SetColumn("CreatedAt", time.Now())
	tx.Statement.SetColumn("CollectedFund", float32(0))

	return nil
}

func (p *Project) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (p *Project) BeforeDelete(tx *gorm.DB) error {
	return errors.New("tidak boleh menghapus data proyek")
}
