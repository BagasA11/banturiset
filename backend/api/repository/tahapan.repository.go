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
	if _, err := tr.GetDataByTahap(t.ProjectID, t.Tahap); err != nil {
		return fmt.Errorf("data tahap %d redundan", t.Tahap)
	}

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

func (tr *TahapRepo) isGT100(projectID uint, input uint) error {
	var p []int
	if err := tr.DB.Model(&models.Tahapan{}).Where("project_id = ?", projectID).Pluck("cost_percent", &p).Error; err != nil {
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
	for v, _ := range in {
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

func (tr *TahapRepo) GetDataByTahap(projectID uint, tahap uint8) (models.Tahapan, error) {
	var t models.Tahapan
	if err := tr.DB.Where("project_id = ? AND tahap = ?", projectID, tahap).First(&t).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return t, nil
		}
		fmt.Printf("getDatabyTahap(): %s/n", err.Error())
		return t, errors.New("gagal mengambil data tahapan")
	}
	return t, errors.New("data redundan")
}
