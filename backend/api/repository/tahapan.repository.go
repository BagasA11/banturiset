package repository

import (
	"errors"

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

func (tr *TahapRepo) Update(t models.Tahapan, penelitiID uint) error {
	var t2 models.Tahapan
	tx := tr.DB.Begin()
	// validasi hak milik
	err := tr.DB.Where("id = ?", t.ID).Joins("Project").First(&t2).Error
	if err != nil {
		return err
	}
	if t2.Project.PenelitiID != penelitiID {
		return errors.New("id peneliti tidak sama")
	}
	if err := tx.Model(&models.Tahapan{}).Where("id = ?", t.ID).Updates(&t).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (tr *TahapRepo) Delete(id uint, penelitiID uint) error {
	var t2 models.Tahapan
	tx := tr.DB.Begin()

	if err := tr.DB.Where("id = ?", id).Joins("Project").First(&t2).Error; err != nil {
		return err
	}
	if t2.Project.PenelitiID != penelitiID {
		return errors.New("id peneliti tidak sama")
	}

	if err := tx.Delete(&models.Tahapan{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
