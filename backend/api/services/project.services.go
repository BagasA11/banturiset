package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/bagasa11/banturiset/api/dto"
	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/api/repository"
	tz "github.com/bagasa11/banturiset/timezone"
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
	now := tz.GetTime(time.Now())
	deadline := now.AddDate(int(req.Year), int(now.Month()), now.Day()) // t->now + year + t->now->month + t->now->day

	fundUntil := now.Add(time.Until(deadline) / 3)
	fmt.Println("mont ", int(deadline.Month()/3))
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
	if err := ps.Repo.SetStatusReject(id); err != nil {
		return err
	}
	if err := ps.Repo.FillRevMsg(id, req.PesanRevisi); err != nil {
		return err
	}

	return nil
}

func (ps *ProjectService) Review(projectID uint) (models.Project, error) {
	return ps.Repo.Review(projectID)
}

func (ps *ProjectService) Verifikasi(projectID uint, adminID uint) (models.Project, error) {
	return ps.Repo.Verifikasi(projectID, adminID)
}

func (ps *ProjectService) UploadProposal(id uint, penelitiID uint, proposalUrl string) error {
	return ps.Repo.UploadProposal(id, penelitiID, proposalUrl)
}

func (ps *ProjectService) UploadKlirens(id uint, penelitiID uint, klirens_url string) error {
	return ps.Repo.UploadKlirens(id, penelitiID, klirens_url)
}

func (ps *ProjectService) Update(id uint, penelitiID uint, req dto.EditProject) error {

	p := models.Project{
		ID:         id,
		Title:      req.Title,
		PenelitiID: penelitiID,
		Desc:       req.Desc,
		Milestone:  req.Milestone,
		TktLevel:   req.Tkt,
		Cost:       req.Cost,
	}

	return ps.Repo.Update(&p)
}

func (ps *ProjectService) Detail(id uint) (models.Project, error) {
	return ps.Repo.Detail(id)
}

func (ps *ProjectService) Preview(id uint, penelitiID uint) (models.Project, error) {
	return ps.Repo.ProjectPreview(id, penelitiID)
}

func (ps *ProjectService) OpenDonate(page uint) ([]models.Project, error) {
	return ps.Repo.OpenDonate(page)
}

func (ps *ProjectService) SubmitToReviewed(ProjectID uint, penelitiID uint) error {
	return ps.Repo.SubmitToReviewed(penelitiID, ProjectID)
}

func IsEditable(id uint) error {
	ps := NewProjectService()
	return ps.Repo.IsEditable(id)
}

func IsOpenFund(id uint) error {
	ps := NewProjectService()
	return ps.Repo.IsOpentoFund(id)
}

func (ps *ProjectService) HasSubmit(page uint) ([]models.Project, error) {
	if page == 0 {
		return []models.Project{}, errors.New("page harus > 0")
	}
	end := page * 20
	start := end - 19

	return ps.Repo.HasSubmit(start, end)
}

func (ps *ProjectService) OnGoing(page uint) ([]models.Project, error) {
	// []models.Project{}
	if page == 0 {
		return []models.Project{}, errors.New("page harus > 0")
	}
	end := page * 20
	start := end - 19

	return ps.Repo.OnGoing(start, end)
}

func (ps *ProjectService) Revisi(penelitiID uint) ([]models.Project, error) {
	return ps.Repo.Revisi(penelitiID)
}

func (ps *ProjectService) MyContributeProject(donaturID uint, page uint) ([]models.Project, error) {
	end := page * 20
	start := end - 19
	return ps.Repo.MyContributeProject(donaturID, start, end)
}

func MyProjectWasClosedDetail(id uint, penelitiID uint, tahap uint8) (models.Project, error) {
	ps := NewProjectService()
	return ps.Repo.MyProjectWasClosedDetail(id, penelitiID, tahap)
}

func (ps *ProjectService) MyProjectCost(projectID uint, penelitiID uint) (float32, error) {
	return ps.Repo.MyProjectCost(projectID, penelitiID)
}
