package services

import (
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

func (bds *BudgetDetailService) Create(projectID uint, penelitiID uint, req dto.BudgetDetailsCreate) error {
	// hitung persentase

	if err := IsMyProject(projectID, penelitiID); err != nil {
		return err
	}

	if err := IsEditable(projectID); err != nil {
		return err
	}

	bd := models.BudgetDetails{
		ProjectID: projectID,
		Deskripsi: req.Desc,
		Cost:      req.Cost,
	}

	return bds.Repo.Create(bd)
}

func (bds *BudgetDetailService) Updates(id uint, req dto.BudgetDetailsCreate, penelitiID uint) error {

	bd := models.BudgetDetails{
		ID:        id,
		Deskripsi: req.Desc,
		Cost:      req.Cost,
	}
	return bds.Repo.Updates(bd, penelitiID)
}

func (bds *BudgetDetailService) Delete(id uint, penelitiID uint) error {
	return bds.Repo.Delete(id, penelitiID)
}
