package services

import (
	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
)

type ProgressServices struct {
	Repo *repository.ReportProgressRepo
}

func NewProgressServices() *ProgressServices {
	return &ProgressServices{
		Repo: repository.NewProgressRepo(),
	}
}

func (ps *ProgressServices) CreateReport(projectID uint, input dto.ProgressReport) error {
	p := models.Progress{
		Tahap:     input.Tahap,
		Desc:      input.Desc,
		FileUrl:   input.FileUrl,
		ProjectID: projectID,
	}
	return ps.Repo.CreateReport(p)
}

func (ps *ProgressServices) IsRedundant(projectID uint, tahap uint8) error {

	return ps.Repo.IsRedundant(projectID, tahap)
}
