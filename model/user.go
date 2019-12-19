package model

import "github.com/jinzhu/gorm"

// User 是一个用户实体
type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"not null"`
}
