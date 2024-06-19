package services

import (
	"errors"
	"time"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
)

type TahapService struct {
	Repo *repository.TahapRepo
}

func NewTahapService() *TahapService {
	return &TahapService{
		Repo: repository.NewTahapRepo(),
	}
}

func (ts *TahapService) Create(projectID uint, penelitiID uint, req dto.TahapCreate) error {

	if err := IsMyProject(projectID, penelitiID); err != nil {
		return err
	}

	if IsEditable(projectID) != nil {
		return errors.New("tidak dapat menambah data tahap pada proyek yang sudah diverifikasi")
	}

	if err := ts.Repo.HasTahap(projectID, uint(req.Tahap)); err != nil {
		return err
	}

	mulai, err := time.Parse(time.RFC3339, req.Start)
	if err != nil {
		return err
	}

	selesai, err := time.Parse(time.RFC3339, req.End)
	if err != nil {
		return err
	}

	t := models.Tahapan{
		ProjectID:   projectID,
		CostPercent: req.CostPercent,
		Tahap:       req.Tahap,
		Start:       mulai,
		End:         selesai,
	}

	return ts.Repo.Create(t)
}

func (ts *TahapService) List(projectID uint, limit uint) ([]models.Tahapan, error) {
	var t []models.Tahapan
	var err error

	if limit == 0 {
		t, err = ts.Repo.All(projectID)
	}

	if limit > 0 {
		t, err = ts.Repo.List(projectID, limit)
	}

	if err != nil {
		return []models.Tahapan{}, err
	}

	return t, nil
}

func (ts *TahapService) Update(id uint, req dto.TahapCreate, projectID uint, penelitiID uint) error {
	mulai, err := time.Parse(time.RFC3339, req.Start)
	if err != nil {
		return err
	}

	selesai, err := time.Parse(time.RFC3339, req.End)
	if err != nil {
		return err
	}

	// memeriksa apakah proyek milik peneliti
	if err = IsMyProject(projectID, penelitiID); err != nil {
		return err
	}
	// memeriksa apakah proyek sudah diverifikasi
	if err = IsEditable(projectID); err != nil {
		return err
	}

	t := models.Tahapan{
		ID:          id,
		ProjectID:   projectID,
		CostPercent: req.CostPercent,
		Start:       mulai,
		End:         selesai,
	}

	return ts.Repo.Update(t)
}

func (ts *TahapService) Delete(id uint, projectID uint, penelitID uint) error {
	if IsMyProject(projectID, penelitID) != nil {
		return errors.New("tidak boleh mengedit/menghapus detail proyek yang bukan milik anda")
	}
	if err := IsEditable(projectID); err != nil {
		return err
	}
	return ts.Repo.Delete(id)
}
