package services

import (
	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
)

type ProgressServices struct {
	Repo *repository.ReportRepo
}

func NewProgressServices() *ProgressServices {
	return &ProgressServices{
		Repo: repository.NewReportRepo(),
	}
}

func (ps *ProgressServices) CreateReport(projectID uint, input dto.ProgressReport) error {
	p := models.Report{
		Desc:      input.Desc,
		FileUrl:   input.FileUrl,
		ProjectID: projectID,
	}
	return ps.Repo.CreateReport(p)
}

// Service untuk memeriksa redundansi laporan
// func (ps *ProgressServices) IsRedundantProgress(projectID uint, tahap uint8) error {
// 	return ps.Repo.IsRedundant(projectID, tahap)
// }
