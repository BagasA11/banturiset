package services

import (
	"time"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
	tz "github.com/bagasa11/banturiset/timezone"
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
	t := tz.GetTime(time.Now()).AddDate(0, int(req.Month), 0)

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

func (ps *PengajuanServices) IsOpen(id uint) error {
	return ps.Repo.IsOpen(id)
}
