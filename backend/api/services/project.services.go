package services

import (
	"time"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
)

type ProjectService struct {
	Repo *repository.ProjectRepository
}

func NewProjectService() *ProjectService {
	return &ProjectService{
		Repo: repository.NewProjectRepository(),
	}
}

func (ps *ProjectService) Create(req dto.CreateProject, penelitiID uint) error {

	deadline := time.Now().AddDate(int(req.Year), int(time.Now().Month()), time.Now().Day())                   // t->now + year + t->now->month + t->now->day
	fundUntil := time.Now().Add(time.Duration((deadline.Year() + int(deadline.Month()) + deadline.Day()) / 3)) // t->now + duration( deadline->y + deadline->m + deadline->y)

	p := models.Project{
		Title:       req.Title,
		Desc:        req.Desc,
		Milestone:   int8(req.Milestone),
		TktLevel:    req.Tkt,
		Cost:        req.Cost,
		DeadLine:    deadline,
		FundUntil:   fundUntil,
		PengajuanID: req.PengajuanID,
		PenelitiID:  penelitiID,
	}

	skema := NewPengajuanService()
	if err := skema.IsOpen(req.PengajuanID); err != nil {
		return err
	}

	return ps.Repo.Create(p)
}

func IsMyProject(id uint, penelitiID uint) error {
	ps := NewProjectService()
	return ps.Repo.IsMyProject(id, penelitiID)
}

func (ps *ProjectService) MyProject(penelitiID uint, limit uint) ([]models.Project, error) {

	return ps.Repo.MyProject(penelitiID, limit)
}

func (ps *ProjectService) Diverifikasi(page uint) ([]models.Project, error) {
	return ps.Repo.Diverifikasi(page)
}

func (ps *ProjectService) Tolak(id uint, req dto.ProjectDitolak) error {
	p := models.Project{
		ID:          id,
		PesanRevisi: &req.PesanRevisi,
		Status:      models.Tolak,
	}
	return ps.Repo.Update(&p)
}

func (ps *ProjectService) Review(projectID uint) (models.Project, error) {
	return ps.Repo.Review(projectID)
}

func (ps *ProjectService) Verifikasi(projectID uint) error {
	return ps.Repo.Verifikasi(projectID)
}

func (ps *ProjectService) UploadProposal(id uint, penelitiID uint, proposalUrl string) error {
	return ps.Repo.UploadProposal(id, penelitiID, proposalUrl)
}

func (ps *ProjectService) UploadKlirens(id uint, penelitiID uint, klirens_url string) error {
	return ps.Repo.UploadKlirens(id, penelitiID, klirens_url)
}

func (ps *ProjectService) Update(id uint, penelitiID uint, req dto.CreateProject) error {
	deadline := time.Now().AddDate(int(req.Year), int(time.Now().Month()), time.Now().Day())

	p := models.Project{
		ID:         id,
		PenelitiID: penelitiID,
		Title:      req.Title,
		Desc:       req.Desc,
		Milestone:  req.Milestone,
		TktLevel:   req.Tkt,
		Cost:       req.Cost,
		DeadLine:   deadline,
	}

	return ps.Repo.Update(&p)
}

func (ps *ProjectService) Detail(id uint) (models.Project, error) {
	return ps.Repo.Detail(id)
}

func (ps *ProjectService) Preview(id uint, penelitiID uint) (models.Project, error) {
	return ps.Repo.ProjectPreview(id, penelitiID)
}

func (ps *ProjectService) SubmitToReviewed(ProjectID uint, penelitiID uint) error {
	return ps.Repo.SubmitToReviewed(penelitiID, ProjectID)
}

func IsEditable(id uint) error {
	ps := NewProjectService()
	return ps.Repo.IsEditable(id)
}
