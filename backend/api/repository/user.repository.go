package repository

import (
	"errors"
	"fmt"

	"github.com/bagasa11/banturiset/api/models"
	"github.com/bagasa11/banturiset/config"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		DB: config.GetDB(),
	}
}

func (ur *UserRepo) Create(u models.User) (uint, error) {
	tx := ur.DB.Begin()
	if err := tx.Create(&u).Error; err != nil {
		tx.Rollback()

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return 0, errors.New("data email atau nomor telepon duplikat")
		}
		return 0, err
	}
	tx.Commit()
	return u.ID, nil
}

func (ur *UserRepo) IsDonatur(userID uint) error {
	var u *models.User
	err := ur.DB.Where("id = ?", userID).Preload("Donatur").Select("id", "role").Limit(1).Find(&u).Error
	if err != nil {
		return err
	}
	if u != nil {
		return fmt.Errorf("user dengan id %d sudah memiliki peran sebagai donatur", userID)
	}
	return nil
}

func (ur *UserRepo) IsPeneliti(userID uint) error {
	var u *models.User
	err := ur.DB.Where("id = ?", userID).Preload("Peneliti").Select("id", "role").Limit(1).Find(&u).Error
	if err != nil {
		return err
	}
	if u != nil {
		return fmt.Errorf("user dengan id %d sudah memiliki peran sebagai peneliti", userID)
	}
	return nil
}

func (ur *UserRepo) FindID(id uint) (models.User, error) {
	var user models.User
	err := ur.DB.Where("id = ?", false).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errors.New("data user tidak ditemukan")
	}
	return user, nil
}

func (ur *UserRepo) CheckID(id uint, role string) error {
	return ur.DB.Select("ID").Where("role = ?", role).First(&models.User{}, id).Error
}

func (ur *UserRepo) SetAvatar(id uint, url string) error {
	tx := ur.DB.Begin()
	if err := tx.Model(&models.User{}).Where("id = ?", id).Update("profile_url", url).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ur *UserRepo) WhereVerified(email string) (models.User, error) {
	var u models.User
	if err := ur.DB.Where("email = ?", email).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (ur *UserRepo) AdminLogin(email string) (models.User, error) {

	var admin models.User
	if err := ur.DB.Where("email = ? AND is_verfied = ? AND is_block = ?", email, true, false).
		Preload("Penyunting").First(&admin).Error; err != nil {
		return admin, err
	}
	return admin, nil
}
func (ur *UserRepo) PenelitLogin(email string) (models.User, error) {
	var peneliti models.User
	if err := ur.DB.Where("email = ? AND is_verfied = ? AND is_block = ?", email, true, false).
		Preload("Peneliti").First(&peneliti).Error; err != nil {
		return peneliti, err
	}
	return peneliti, nil
}

func (ur *UserRepo) DonaturLogin(email string) (models.User, error) {
	var donatur models.User
	if err := ur.DB.Where("email = ? AND is_verfied = ? AND is_block = ?", email, true, false).
		Preload("Donatur").First(&donatur).Error; err != nil {
		return donatur, err
	}
	return donatur, nil

}

func (ur *UserRepo) DonaturProfile(userID uint) (*models.User, uint, error) {
	u := models.User{}
	if err := ur.DB.Where("id = ? AND is_verfied = ? AND is_block = ?", userID, true, false).Joins("Donatur").First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.New("data user tidak ditemukan")
		}
		return nil, 0, err
	}
	return &u, u.Donatur.ID, nil
}

func (ur *UserRepo) PenelitiProfile(userID uint) (*models.User, uint, error) {
	u := models.User{}
	if err := ur.DB.Where("id = ? AND is_verfied = ? AND is_block = ?", userID, true, false).Joins("Peneliti").First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.New("data user tidak ditemukan")
		}
		return nil, 0, err
	}
	return &u, u.Peneliti.ID, nil
}

func (ur *UserRepo) AdminProfile(userID uint) (*models.User, uint, error) {
	u := models.User{}
	if err := ur.DB.Where("id = ? AND is_verfied = ? AND is_block = ?", userID, true, false).Joins("Penyunting").First(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, 0, errors.New("data user tidak ditemukan")
		}
		return nil, 0, err
	}
	return &u, u.Penyunting.ID, nil
}

func (ur *UserRepo) EmailLogin(email string) (*models.User, error) {
	u := new(models.User)
	if err := ur.DB.Where("email = ?", email).First(&u).Error; err != nil {
		fmt.Println("error ur->emailLgin(): ", err.Error())
		return nil, errors.New("login gagal")
	}
	return u, nil
}

func (pr *UserRepo) Update(user *models.User) error {
	if user == nil {
		return errors.New("informasi update kosong")
	}
	tx := pr.DB.Begin()
	if err := tx.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (ur *UserRepo) GetPayment(userID uint) (models.PaymentInfo, error) {
	p := models.PaymentInfo{}
	err := ur.DB.Table("users").Where("id = ?", userID).Scan(&p).Error
	return p, err
}
