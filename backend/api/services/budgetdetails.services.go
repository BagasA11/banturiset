package services

import (
	"errors"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
)

type BudgetDetailService struct {
	Repo *repository.BudgetDetailsRepo
}

func NewBudgetDetailService() *BudgetDetailService {
	return &BudgetDetailService{
		Repo: repository.NewBudgetDetailsRepo(),
	}
}

func (bds *BudgetDetailService) Create(projectID uint, req dto.BudgetDetailsCreate) error {
	// hitung persentase
	persent, err := bds.Repo.Percentage(projectID)
	if err != nil {
		return err
	}
	if persent >= 100 {
		return errors.New("persentase alokasi pendanaan sudah mencapai 100%")
	}

	bd := models.BudgetDetails{
		ProjectID: projectID,
		Deskripsi: req.Desc,
		Cost:      req.Cost,
		Tahap:     req.Tahap,
		Percent:   req.Percent,
	}

	return bds.Repo.Create(bd)
}

func (bds *BudgetDetailService) GetPercent(projectID uint) (int8, error) {
	return bds.Repo.Percentage(projectID)
}

func (bds *BudgetDetailService) Updates(id uint, req dto.BudgetDetailsCreate) error {

	bd := models.BudgetDetails{
		ID:        id,
		Deskripsi: req.Desc,
		Cost:      req.Cost,
		Tahap:     req.Tahap,
		Percent:   req.Percent,
	}
	return bds.Repo.Updates(bd)
}

func (bds *BudgetDetailService) Delete(id uint) error {
	return bds.Repo.Delete(id)
}
