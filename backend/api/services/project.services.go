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

func (ps *ProjectService) Create(req dto.ProjectCreate) error {
	dl, err := time.Parse(time.RFC3339, req.DeadLine)
	if err != nil {
		return err
	}
	p := models.Project{
		Title:     req.Title,
		Desc:      req.Desc,
		Milestone: int8(req.Milestone),
		TktLevel:  req.Tkt,
		Cost:      req.Cost,
		DeadLine:  dl,
	}
	return ps.Repo.Create(p)
}

func (ps *ProjectService) MyProject(penelitiID uint, limit uint) ([]models.Project, error) {

	return ps.Repo.MyProject(penelitiID, limit)
}

func (ps *ProjectService) Diverifikasi(page uint) ([]models.Project, error) {
	return ps.Repo.Diverifikasi(page)
}

func (ps *ProjectService) Ditolak(id uint, req dto.ProjectDitolak) error {
	p := models.Project{
		ID:          id,
		PesanRevisi: &req.PesanRevisi,
		Status:      repository.Tolak,
	}
	return ps.Repo.Update(&p)
}

func (ps *ProjectService) Review(projectID uint) (models.Project, error) {
	return ps.Repo.Review(projectID)
}
