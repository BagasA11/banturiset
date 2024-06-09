package repository

import (
	"errors"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

const Fraud = true
const Abort = -2
const Tolak = -1
const Draft = 0
const Verifikasi = 1
const Selesai = 2

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

	if p.Pengajuan.ClosedAt.Nanosecond() < p.CreatedAt.Nanosecond() {
		tx.Rollback()
		return errors.New("waktu pengajuan sudah ditutup")
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

func (pr *ProjectRepository) IsMyProject(id uint, penelitiID uint) error {
	return pr.DB.Where("peneliti_id = ?", penelitiID).First(&models.Project{}, id).Error
}

func (pr *ProjectRepository) Diverifikasi(page uint) ([]models.Project, error) {
	var ps []models.Project
	// page 1 : 1 - 10
	// page 2 : 11 - 20
	// page 3 : 21 - 30

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
	// []tahapan{}

	var p models.Project
	if err := pr.DB.Where("id = ? AND status BETWEEN ? AND ?", id, Tolak, Draft).Where("fraud = ?", false).Preload("BudgetDetails").
		Preload("Tahapan").Joins("Pengajuan").First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (pr *ProjectRepository) Detail(id uint) (models.Project, error) {
	// project{}
	// []budget{}
	// []tahapan{}

	var p models.Project
	if err := pr.DB.Where("id = ? AND status >= ? AND fraud = ?", id, Verifikasi, false).Preload("BudgetDetails").
		Preload("Tahapan").Joins("Pengajuan").First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (pr *ProjectRepository) Verifikasi(id uint) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND fraud = ?", id, !Fraud).
		Where("klirens_url IS NOT NULL AND proposal_url IS NOT NULL").
		Update("status", Verifikasi).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) Update(p *models.Project) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND peneliti_id = ? AND status < ?", p.ID, p.Peneliti.ID, Verifikasi).Updates(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) UploadKlirens(id uint, penelitiID uint, klirens_url string) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND peneliti_id = ? AND status < ?", id, penelitiID, Verifikasi).Update("klirens_url", klirens_url).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) UploadProposal(id uint, penelitiID uint, proposalUrl string) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND peneliti_id = ? AND status < ?", id, penelitiID, Verifikasi).
		Update("proposal_url", proposalUrl).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) Abort(id uint)   {}
func (pr *ProjectRepository) Selesai(id uint) {}
