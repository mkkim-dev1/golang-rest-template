package db

import (
	"serengeti.app/go-rest-template/pkg/config"
	"serengeti.app/go-rest-template/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and migrates models
func InitDB() error {
	var err error
	if config.AppConfig.DbKind == "sqlite" {
		DB, err = gorm.Open(sqlite.Open(config.AppConfig.DSN), &gorm.Config{})
	} else if config.AppConfig.DbKind == "mysql" || config.AppConfig.DbKind == "maria" {
		DB, err = gorm.Open(mysql.Open(config.AppConfig.DSN), &gorm.Config{})
	}
	if err != nil {
		return err
	}
	return DB.AutoMigrate(&models.Book{}) // Book 모델 마이그레이션
}
