package repository

import (
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

func (bdr *BudgetDetailsRepo) Percentage(projectID uint) (int8, error) {
	var sum int64
	err := bdr.DB.Select("SUM(percent)").Where("project_id = ?", projectID).Table("budget_details").Error
	return int8(sum), err
}

func (bdr *BudgetDetailsRepo) CountRows(projectID uint) (uint8, error) {
	var count int64
	if err := bdr.DB.Model(&models.BudgetDetails{}).Where("project_id = ?", projectID).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint8(count), nil
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

func (bdr *BudgetDetailsRepo) Updates(bd models.BudgetDetails) error {
	tx := bdr.DB.Begin()
	if err := tx.Model(&models.BudgetDetails{}).Where("id = ?", bd.ID).Updates(&bd).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (bdr *BudgetDetailsRepo) Delete(id uint) error {
	tx := bdr.DB.Begin()
	if err := tx.Where("id = ?", id).Delete(&models.BudgetDetails{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
