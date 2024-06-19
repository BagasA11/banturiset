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

func (bds *BudgetDetailService) Updates(bdID uint, req dto.BudgetDetailsCreate, projectID uint, penelitiID uint) error {

	if err := IsMyProject(projectID, penelitiID); err != nil {
		return err
	}
	if err := IsEditable(projectID); err != nil {
		return err
	}

	bd := models.BudgetDetails{
		ID:        bdID,
		Deskripsi: req.Desc,
		Cost:      req.Cost,
	}
	return bds.Repo.Updates(bd)
}

func (bds *BudgetDetailService) Delete(id uint, projectID uint, penelitiID uint) error {
	if IsMyProject(projectID, penelitiID) != nil {
		return errors.New("tidak diperbolehkan mengubah/menghapus item proyek yang bukan milik anda")
	}
	if err := IsEditable(projectID); err != nil {
		return err
	}

	return bds.Repo.Delete(id)
}
