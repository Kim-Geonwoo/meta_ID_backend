package models

import "gorm.io/gorm"

// User 모델 정의
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
