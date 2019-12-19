package model

import (
	"github.com/jinzhu/gorm"
)

// DB 是gorm这个包里一个连接数据库的指针，我们可以通过这个指针来对数据库进行操作
var DB *gorm.DB

// Init 初始化admin:123456
func Init(connString string) error {
	// connString是一个字符串，他保存着连接数据所需要的信息
	db, err := gorm.Open("mysql", connString)
	// db 就是这个函数返回的一个连接mysql的一个实例
	if err != nil {
		panic(err)
	}
	DB = db
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{})
	// 创建用户
	DB.Create(&User{
		Username: "admin",
		Password: "123456",
	})
	return nil
}
