package models

// import (
// 	"errors"
// 	"slices"
// 	"strings"
// 	"time"

// 	"gorm.io/gorm"
// )

// type Project struct {
// 	ID          uint   `gorm:"primaryKey"`
// 	Title       string `gorm:"size:50"`
// 	Desc        *string
// 	ProposalUrl *string `gorm:"size:170"`
// 	KlirensUrl  *string `gorm:"size:170"`
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// 	FundUntil   time.Time
// 	DeadLine    time.Time
// 	Years       int8 `gorm:"not null; default:1"`
// 	Milestone   int8 `gorm:"not null; default:1"`
// 	TktLevel    int8 `gorm:"not null; default:1"`
// 	Cost        float32
// 	Status      string `gorm:"not null; size:10 "`

// 	PenelitiID uint
// 	KategoriID uint

// 	Kategori      Kategori `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; embedded;"`
// 	Peneliti      Peneliti `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; embedded"`
// 	BudgetDetails []BudgetDetails
// }

// func (u *Project) BeforeCreate(tx *gorm.DB) error {
// 	tx.Statement.SetColumn("Status", "draft")
// 	tx.Statement.SetColumn("CreatedAt", time.Now())
// 	tx.Statement.SetColumn("DeadLine", (u.CreatedAt.Add(time.Duration(u.Years))))
// 	// if fund > created_at + 3 bulan : error
// 	// created_at + 1h * 24h * 30d * 3month => created_at + 3 months
// 	if u.FundUntil.After(u.CreatedAt.Add(time.Hour * 24 * 30 * 4)) {
// 		return errors.New("durasi pendanaan maksimum 3 bulan")
// 	}
// 	return nil
// }

// func (p *Project) BeforeUpdate(tx *gorm.DB) error {
// 	status := []string{"diverifikasi", "terverifikasi", "verified", "verify", "valid", "validated", "divalidasi",
// 		"accepted", "diterima"}
// 	if slices.Contains(status, strings.ToLower(p.Status)) {
// 		return errors.New("proyek yang sudah diverifikasi tidak dapat diubah")
// 	}
// 	tx.Statement.SetColumn("UpdatedAt", time.Now())
// 	return nil
// }
