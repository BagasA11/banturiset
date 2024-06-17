package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Tahapan struct {
	ID          uint      `gorm:"primaryKey"`
	CostPercent uint8     `gorm:"not null; default:0"`
	Tahap       uint8     `gorm:"not null; default:0"`
	Start       time.Time `gorm:"not null"`
	End         time.Time `gorm:"not null"`

	ProjectID uint
	Project   Project
}

func (t *Tahapan) BeforeCreate(tx *gorm.DB) error {
	if t.CostPercent > 100 {
		return errors.New("persentase harus < 100")
	}
	if t.End.Nanosecond() < time.Now().Nanosecond() {
		return errors.New("deadline tahapan harus lebih dari masa kini")
	}

	if t.Start.After(t.End) {
		return fmt.Errorf("waktu mulai tahapan: %v tidak boleh mendahului deadline %v", t.Start, t.End)
	}
	return nil
}

func (t *Tahapan) BeforeUpdate(tx *gorm.DB) error {
	if t.Project.Status >= Verifikasi {
		return errors.New("proyek sudah disetujui. tidak dapat mengedit tahapan")
	}
	return nil
}

func (t *Tahapan) BeforeDelete(tx *gorm.DB) error {
	if t.Project.Status >= Verifikasi {
		return errors.New("proyek sudah disetujui. tidak dapat mengedit tahapan")
	}
	return nil
}
