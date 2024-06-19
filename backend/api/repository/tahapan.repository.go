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
	if err := tx.Model(&models.Tahapan{}).Where("id = ?", t.ID).Updates(&t).Error; err != nil {
		tx.Rollback()
		fmt.Println("[repo] tahap->Update(): ", err.Error())
		return errors.New("gagal mengubah tahap")
	}
	tx.Commit()
	return nil
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
