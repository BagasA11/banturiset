package models

import (
	"errors"

	"time"

	tz "github.com/bagasa11/banturiset/timezone"
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
	Status        int8      `gorm:"not null; size:10;"`
	PesanRevisi   *string
	Fraud         bool    `gorm:"not null; default:false"`
	CollectedFund float32 `gorm:"not null"`

	PengajuanID uint      `gorm:"not null"`
	Pengajuan   Pengajuan `gorm:"foreignKey:PengajuanID"`
	PenelitiID  uint      `gorm:"not null"`
	Peneliti    Peneliti  `gorm:"foreignKey:PenelitiID"`
	AdminID     *uint
	Penyunting  Penyunting `gorm:"foreignKey:AdminID"`

	Donasi        []Donasi
	BudgetDetails []BudgetDetails
	Tahapan       []Tahapan
	Payout        []Payout
	Progress      []Progress
}

func (p *Project) BeforeCreate(tx *gorm.DB) error {

	if p.Cost < float32(0) {
		return errors.New("biaya harus > 0")
	}
	if p.Milestone < 1 {
		return errors.New("milestones minimum adalah 1")
	}
	if p.TktLevel < 1 {
		return errors.New("level tkt minimum adalah 1")
	}

	tx.Statement.SetColumn("CreatedAt", tz.GetTime(time.Now()))
	tx.Statement.SetColumn("Status", Draft)
	tx.Statement.SetColumn("CollectedFund", float32(0))

	return nil
}

func (p *Project) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", tz.GetTime(time.Now()))
	return nil
}

func (p *Project) BeforeDelete(tx *gorm.DB) error {
	return errors.New("tidak boleh menghapus data proyek")
}
