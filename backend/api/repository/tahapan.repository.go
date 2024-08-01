package repository

import (
	"errors"
	"fmt"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

// const (
// 	tolak  int8 = -1
// 	draft  int8 = 0
// 	terima int8 = 0
// )

type TahapRepo struct {
	DB *gorm.DB
}

func NewTahapRepo() *TahapRepo {
	return &TahapRepo{
		DB: config.GetDB(),
	}
}

func (tr *TahapRepo) Create(t models.Tahapan) error {
	// ensure cost not more than 100%
	if err := tr.isGT100(t.ProjectID, uint(t.CostPercent)); err != nil {
		return err
	}
	// if _, err := tr.GetDataByTahap(t.ProjectID, t.Tahap); err != nil {
	// 	return fmt.Errorf("data tahap %d redundan", t.Tahap)
	// }

	tx := tr.DB.Begin()
	if err := tx.Create(&t).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (tr *TahapRepo) All(projectID uint) ([]models.Tahapan, error) {
	var t []models.Tahapan
	err := tr.DB.Where("project_id = ?", projectID).Find(&t).Error
	return t, err
}

func (tr *TahapRepo) List(projectID uint, limit uint) ([]models.Tahapan, error) {
	var t []models.Tahapan
	err := tr.DB.Where("project_id = ?", projectID).Limit(int(limit)).Find(&t).Error
	return t, err
}

func (tr *TahapRepo) Update(t models.Tahapan) error {
	tx := tr.DB.Begin()
	if err := tr.isGT100(t.ProjectID, uint(t.CostPercent)); err != nil {
		return err
	}

	if err := tx.Model(&models.Tahapan{}).Where("id = ?", t.ID).Updates(&t).Error; err != nil {
		tx.Rollback()
		fmt.Println("[repo] tahap->Update(): ", err.Error())
		return errors.New("gagal mengubah tahap")
	}
	tx.Commit()
	return nil
}

// type percent struct {
// 	percent uint
// }

func (tr *TahapRepo) GetTahap(id uint, project_id uint) (uint8, error) {
	var tahap []uint8
	if err := tr.DB.Model(&models.Tahapan{}).Where("id = ? AND project_id = ?", id, project_id).Limit(1).Pluck("tahap", &tahap).Error; err != nil {
		return 0, err
	}
	if len(tahap) == 0 {
		return 0, errors.New("gagal mengambil data tahapan atau data tidak ada")
	}
	fmt.Println("tahap->pluck():\t", tahap[0])
	return tahap[0], nil
}

func (tr *TahapRepo) isGT100(projectID uint, input uint) error {
	var p []int
	if err := tr.DB.
		Model(&models.Tahapan{}).Where("project_id = ?", projectID).
		Pluck("cost_percent", &p).Error; err != nil {
		fmt.Printf("isGT100(): %s\n", err.Error())
		return err
	}
	if sum(p)+int(input) > 100 {
		return errors.New("persentase tidak boleh melebihi 100%")
	}
	return nil
}

func sum(in []int) int {
	var sum int = 0
	for _, v := range in {
		sum += v
	}
	return sum
}

func (tr *TahapRepo) Delete(id uint) error {

	tx := tr.DB.Begin()

	if err := tx.Delete(&models.Tahapan{}, id).Error; err != nil {
		tx.Rollback()
		fmt.Println("[repo] tahapan->delete(): ", err.Error())
		return errors.New("gagal menghapus tahap")
	}
	tx.Commit()
	return nil
}

func (tr *TahapRepo) IsTahapRedundant(projectID uint, tahap uint8) error {
	var t models.Tahapan
	if err := tr.DB.Where("project_id = ? AND tahap = ?", projectID, tahap).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		fmt.Printf("getDatabyTahap(): %s/n", err.Error())
		return errors.New("gagal mengambil data tahapan")
	}
	return errors.New("data redundan")
}

// suatu tahap tidak boleh dibuat atau dihapus ketika mememiliki suksesor. ex: membuat tahap ke-1 ketika ada tahap ke-2
func (tr *TahapRepo) HasSuccessor(projectID uint, tahap uint8) error {
	t := new(models.Tahapan)
	err := tr.DB.Where("project_id = ? AND tahap = ?", projectID, tahap+1).First(&t).Error
	fmt.Println(t.ID, "\t", t.Tahap)
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if t != nil {
		return errors.New("tahap tidak boleh memiliki suksesor")
	}
	return errors.New("gagal memvalidasi tahap suksesor")
}

// tahapan harus memiliki pendahulu
//
// ex: sebelum menambahkan data tahapan ke-2, maka harus ada tahapan ke-1
func (tr *TahapRepo) HasNotAncestor(projectID uint, tahap uint8) error {

	if tahap == 1 {
		return nil
	}
	err := tr.DB.Where("project_id = ? AND tahap = ?", projectID, tahap-1).First(&models.Tahapan{}).Error
	if err == gorm.ErrRecordNotFound {
		return gorm.ErrRecordNotFound
	}
	return err
}

func (tr *TahapRepo) GetTahapList(projectID uint) ([]uint8, error) {
	var tahap []uint8
	if err := tr.DB.Model(&models.Tahapan{}).Where("project_id = ?", projectID).Pluck("tahap", &tahap).Error; err != nil {
		return []uint8{}, err
	}
	return tahap, nil
}
