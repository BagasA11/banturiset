package config

import (
	"fmt"
	"os"

	"github.com/bagasa11/banturiset/api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error when connect to database: ", err.Error())
		return err
	}

	err = DB.AutoMigrate(&models.User{}, &models.Donatur{}, &models.Peneliti{}, &models.Penyunting{}, &models.Kategori{})
	// err = DB.AutoMigrate(&models.User{}, &models.Donatur{}, &models.Peneliti{}, &models.Penyunting{}, &models.Project{},
	// 	&models.Kategori{}, &models.BudgetDetails{})
	if err != nil {
		fmt.Println("error when running migration: ", err.Error())
	}

	return err
}

func GetDB() *gorm.DB {
	return DB
}
