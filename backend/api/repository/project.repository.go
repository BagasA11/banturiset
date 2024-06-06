package repository

import (
	"errors"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

const Fraud = true
const Tolak = -1
const draft = 0
const Verifikasi = 1

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository() *ProjectRepository {
	return &ProjectRepository{
		DB: config.GetDB(),
	}
}

func (pr *ProjectRepository) Create(p models.Project) error {
	tx := pr.DB.Begin()
	if err := tx.Create(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) MyProject(penelitiID uint, limit uint) ([]models.Project, error) {
	var ps []models.Project

	if limit == 0 {
		return ps, errors.New("limit item harus > 0")
	}

	err := pr.DB.Where("peneliti_id = ?", penelitiID).Limit(int(limit)).Find(&ps).Error
	return ps, err
}

func (pr *ProjectRepository) Diverifikasi(page uint) ([]models.Project, error) {
	var ps []models.Project
	// page 1 : 1 - 10
	// page 2 : 11 - 20
	// page 3 : 21 - 30
	// []trans{}
	if page == 0 {
		return ps, errors.New("page harus diatas 0")
	}
	var to = page * 10

	if err := pr.DB.Where("id BETWEEN ? AND ?", to-9, to).Where("status >= ?", Verifikasi).Find(&ps).Error; err != nil {
		return ps, err
	}
	return ps, nil
}

func (pr *ProjectRepository) Review(id uint) (models.Project, error) {
	// project{}
	// []budget{}

	var p models.Project
	if err := pr.DB.Where("id = ? AND status = ?", id, draft).Preload("BudgetDetails").Joins("Pengajuan").First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (pr *ProjectRepository) Verifikasi(id uint) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND fraud = ?", id, !Fraud).Update("status", Verifikasi).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) Update(p *models.Project) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ?", p.ID).Updates(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
