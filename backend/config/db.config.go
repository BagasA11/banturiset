package config

import (
	"fmt"
	"os"

	"github.com/bagasa11/banturiset/api/models"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func getMysqlDsn() string {
	return fmt.Sprintf("%s:%s@tcp(host.docker.internal:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASS"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
}

func get_postgresDsn() string {
	return fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("POSTGRES_USERNAME"), os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_DBName"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_TIMEZONE"))
}

func InitDB() error {
	dsn := getMysqlDsn()

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error when connect to database: ", err.Error())
		return err
	}

	err = DB.AutoMigrate(&models.User{}, &models.Donatur{}, &models.Peneliti{},
		&models.Penyunting{}, &models.Pengajuan{}, &models.Project{}, &models.BudgetDetails{},
		&models.Tahapan{}, &models.Progress{}, &models.Donasi{}, &models.TokenList{}, &models.Payout{})

	// err = DB.AutoMigrate(&models.User{}, &models.Donatur{}, &models.Peneliti{},
	// 	&models.Penyunting{}, &models.Pengajuan{}, &models.Project{}, &models.BudgetDetails{},
	// 	&models.Tahapan{}, &models.Progress{}, &models.Donasi{})

	if err != nil {
		fmt.Println("error when running migration: ", err.Error())
	}

	return err
}

func GetDB() *gorm.DB {
	return DB
}
