package db

import (
	"serengeti.app/go-rest-template/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and migrates models
func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&models.Book{}) // Book 모델 마이그레이션
}
