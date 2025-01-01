package database

import (
	"log"
	"meta_ID_backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 데이터베이스 마이그레이션
	DB.AutoMigrate(&models.User{})
}
