package config

import (
	"cv-builder-api/internal/models"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *Config) {
	// load file .env
	err := godotenv.Load()
	if err != nil {
		panic("Gagal Memuat File .Env")
	}

	// ambil data .env
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Gagal koneksi: %v", err))
	}

	database.AutoMigrate(&models.User{},
		&models.Profile{},
		&models.WorkExperience{},
		&models.Education{},
		&models.Skills{},
		&models.Projects{})

	DB = database
	fmt.Println("Koneksi datbase berhasil")
}
