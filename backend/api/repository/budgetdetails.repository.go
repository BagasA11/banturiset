package repository

import (
	"errors"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type BudgetDetailsRepo struct {
	DB *gorm.DB
}

func NewBudgetDetailsRepo() *BudgetDetailsRepo {
	return &BudgetDetailsRepo{
		DB: config.GetDB(),
	}
}

func (bdr *BudgetDetailsRepo) Create(bd models.BudgetDetails) error {
	tx := bdr.DB.Begin()
	if err := tx.Create(&bd).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// func (bdr *BudgetDetailsRepo) Sum(projectID uint) (float32, error) {
// 	var sum float32
// 	if err := bdr.DB.Where("project_id = ?", projectID).Select("SELECT SUM(cost) FROM budget_details").Row().Scan(&sum); err != nil {
// 		return 0, err
// 	}
// 	return float32(sum), nil
// }

func (bdr *BudgetDetailsRepo) Updates(bd models.BudgetDetails, penelitiID uint) error {
	bd2 := models.BudgetDetails{}
	tx := bdr.DB.Begin()

	if err := bdr.DB.Where("id = ?", bd.ID).Joins("Project").First(&bd2).Error; err != nil {
		return err
	}
	if bd2.Project.PenelitiID != penelitiID {
		return errors.New("id peneliti tidak sama")
	}

	if bd2.Project.Status >= models.Verifikasi {
		return errors.New("tidak dapat mengedit detail budget pada proyek yang sudah diverifikasi")
	}

	if err := tx.Model(&models.BudgetDetails{}).Where("id = ?", bd.ID).Updates(&bd).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (bdr *BudgetDetailsRepo) Delete(id uint, penelitiID uint) error {
	bd := models.BudgetDetails{}
	tx := bdr.DB.Begin()

	if err := bdr.DB.Where("id = ?", id).Joins("Project").First(&bd).Error; err != nil {
		return err
	}
	if bd.Project.PenelitiID != penelitiID {
		return errors.New("id peneliti tidak sama")
	}

	if bd.Project.Status >= models.Verifikasi {
		return errors.New("tidak dapat menghapus detail budget pada proyek yang sudah diverifikasi")
	}

	if err := tx.Delete(&models.BudgetDetails{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
