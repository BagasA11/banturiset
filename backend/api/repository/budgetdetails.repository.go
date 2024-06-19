package repository

import (
	"errors"
	"fmt"

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

func (bdr *BudgetDetailsRepo) Updates(bd models.BudgetDetails) error {

	tx := bdr.DB.Begin()

	if err := tx.Model(&models.BudgetDetails{}).Where("id = ?", bd.ID).Updates(&bd).Error; err != nil {
		tx.Rollback()
		fmt.Println("[repo] budget_details->update: ", err.Error())
		return errors.New("gagal mengupdate detail budget")
	}

	tx.Commit()
	return nil
}

func (bdr *BudgetDetailsRepo) Delete(id uint) error {

	tx := bdr.DB.Begin()
	if err := tx.Delete(&models.BudgetDetails{}, id).Error; err != nil {
		tx.Rollback()
		fmt.Println("[repo] BudgetDetails->delete(): ", err.Error())
		return errors.New("gagal menghapus budget detail")
	}

	tx.Commit()
	return nil
}
