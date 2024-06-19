package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

// const Fraud = true
// const Abort = -2
// const Tolak = -1
// const Draft = 0
// const Verifikasi = 1
// const Selesai = 2

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
		fmt.Printf("error project->Create(): %v", err)
		return errors.New("gagal membuat data project")
	}

	// if p.Pengajuan.ClosedAt.Nanosecond() < p.CreatedAt.Nanosecond() {
	// 	tx.Rollback()
	// 	return errors.New("waktu pengajuan sudah ditutup")
	// }

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
	return pr.DB.Where("peneliti_id = ? AND id = ?", penelitiID, id).First(&models.Project{}).Error
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

	if err := pr.DB.Where("id BETWEEN ? AND ?", to-9, to).Where("status >= ?", models.Verifikasi).Find(&ps).Error; err != nil {
		fmt.Printf("\nerror project->diverifikasi(): %s", err.Error())
		return ps, errors.New("gagal mendapatkan daftar proyek")
	}
	return ps, nil
}

func (pr *ProjectRepository) HasSubmit(start uint, end uint) ([]models.Project, error) {
	var p []models.Project
	if err := pr.DB.Where("status = ? AND fraud = ? AND id BETWEEN ? AND ?", models.Submit, !models.Fraud, start, end).
		Find(&p).Error; err != nil {
		fmt.Println("[repo] project->hasSubmit(), ", err.Error())
		return []models.Project{}, errors.New("gagal mengambil data")
	}
	return p, nil
}

func (pr *ProjectRepository) ProjectPreview(id uint, penelitiID uint) (models.Project, error) {
	var p models.Project
	if err := pr.DB.Where("peneliti_id = ? AND status > ? AND status <= ? AND fraud = ?",
		penelitiID, models.Abort, models.Draft, false).Preload("BudgetDetails").Preload("Tahapan").Find(&p, id).
		Error; err != nil {

		fmt.Println("error [repo] project->preview(): ", err.Error())
		return models.Project{}, errors.New("gagal mendapatkan detail project atau project ini telah diverifikasi")
	}
	return p, nil
}

func (pr *ProjectRepository) Review(id uint) (models.Project, error) {
	// project{}
	// []budget{}
	// []tahapan{}

	var p models.Project
	err := pr.DB.Where("status BETWEEN ? AND ? AND fraud = ?", models.Tolak, models.Submit, !models.Fraud).
		Preload("Pengajuan").Preload("Tahapan").Preload("BudgetDetails").Limit(1).Find(&p, id).Error
	if err != nil {
		fmt.Println("error project->Review(): ", err.Error())
		return models.Project{}, errors.New("gagal mereview proyek")
	}

	return p, nil
}

func (pr *ProjectRepository) Detail(id uint) (models.Project, error) {
	// project{}
	// []budget{}
	// []tahapan{}

	var p models.Project
	if err := pr.DB.Where("id = ? AND status >= ? AND fraud = ?", id, models.Verifikasi, false).Preload("BudgetDetails").
		Preload("Tahapan").Joins("Pengajuan").First(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}

func (pr *ProjectRepository) Verifikasi(id uint) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND fraud = ?", id, !models.Fraud).
		Where("klirens_url IS NOT NULL AND proposal_url IS NOT NULL").
		Update("status", models.Verifikasi).Error; err != nil {
		tx.Rollback()
		fmt.Println("\n error project->verifikasi(): ", err.Error())
		return errors.New("gagal memverifikasi proyek")
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) Update(p *models.Project) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND peneliti_id = ? AND status < ?", p.ID, p.Peneliti.ID, models.Submit).Updates(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) UploadKlirens(id uint, penelitiID uint, klirens_url string) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND peneliti_id = ? AND status < ?", id, penelitiID, models.Submit).Update("klirens_url", klirens_url).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) UploadProposal(id uint, penelitiID uint, proposalUrl string) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND peneliti_id = ? AND status < ?", id, penelitiID, models.Submit).
		Update("proposal_url", proposalUrl).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) TambahSaldo(id uint, saldo float32) error {
	tx := pr.DB.Begin()
	p := models.Project{}
	if err := pr.DB.First(&p, id).Error; err != nil {
		return err
	}

	p.CollectedFund += saldo
	if err := tx.Save(&p).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) OpenDonate(page uint) ([]models.Project, error) {
	var p []models.Project
	if page == 0 {
		return []models.Project{}, errors.New("page tidak boleh kosong")
	}
	pAkhir := uint(page * 10)
	pAwal := uint(pAkhir - 9)

	err := pr.DB.
		Where("status >= ? AND fraud = ? AND id BETWEEN ? AND ?", models.Verifikasi, !models.Fraud, pAwal, pAkhir).
		Where("fund_until >= ?", time.Now()).
		Find(&p).Error
	if err != nil {
		fmt.Println(err)
	}
	return p, err
}

func (pr *ProjectRepository) OnGoing(start uint, end uint) ([]models.Project, error) {
	var p []models.Project
	if err := pr.DB.
		Where("id BETWEEN ? AND ? AND status >= ? AND fund_until < ?", start, end, models.Verifikasi, time.Now()).Find(&p).Error; err != nil {
		fmt.Println("[repo] project->ongoing()", err.Error())
		return p, errors.New("gagal mengambil data")
	}

	return p, nil
}

func (pr *ProjectRepository) IsOpentoFund(id uint) error {

	// mencari proyek yang sudah diverifikasi
	if err := pr.DB.
		Where("status >= ?", models.Verifikasi).Where("fund_until >= ?", time.Now()).First(&models.Project{}, id).Error; err != nil {
		return errors.New("gagal mengecek timeline pendanaan proyek")
	}

	return nil
}

func (pr *ProjectRepository) IsEditable(id uint) error {

	if err := pr.DB.Where("status < ? AND id = ?", models.Submit, id).First(&models.Project{}).Error; err != nil {
		return errors.New("proyek tidak bisa diubah")
	}
	return nil
}

func (pr *ProjectRepository) SubmitToReviewed(penelitiID uint, projectID uint) error {
	tx := pr.DB.Begin()
	if err := tx.Model(&models.Project{}).Where("id = ? AND peneliti_id = ?", projectID, penelitiID).
		Update("status", models.Submit).Error; err != nil {
		tx.Rollback()
		fmt.Println("[repo] project->submitToReviewed(): ", err.Error())
		return errors.New("gagal mensubmit proyek")
	}
	tx.Commit()
	return nil
}

func (pr *ProjectRepository) Revisi(penelitiID uint) ([]models.Project, error) {
	p := []models.Project{}
	if err := pr.DB.Where("peneliti_id = ? AND status = ? AND fraud = ?", penelitiID, models.Tolak, !models.Fraud).
		Find(&p).Error; err != nil {
		fmt.Println("[repo] project->revisi(): ", err.Error())
		return p, errors.New("gagal mengambil data proyek")
	}
	return p, nil
}

func (pr *ProjectRepository) Abort(id uint)   {}
func (pr *ProjectRepository) Selesai(id uint) {}
