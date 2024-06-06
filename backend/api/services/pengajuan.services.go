package services

import (
	"time"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
)

type PengajuanServices struct {
	Repo *repository.PengajuanRepository
}

func NewPengajuanService() *PengajuanServices {
	return &PengajuanServices{
		Repo: repository.NewPengajuanRepository(),
	}
}

func (ps *PengajuanServices) Create(req dto.Pengajuan, penyuntingID uint) error {
	t, err := time.Parse(time.RFC3339, req.ClosedAt)
	if err != nil {
		return err
	}

	p := models.Pengajuan{
		Title:        req.Title,
		Desc:         req.Desc,
		LinkWa:       req.LinkWA,
		LinkPanduan:  req.LinkPanduan,
		IconUrl:      req.LinkIcon,
		ClosedAt:     t,
		PenyuntingID: penyuntingID,
	}

	return repository.NewPengajuanRepository().Create(p)
}

func (ps *PengajuanServices) Open() ([]models.Pengajuan, error) {
	return ps.Repo.Open()
}
