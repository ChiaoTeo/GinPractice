package conf

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// DB 是gorm这个包里一个连接数据库的指针，我们可以通过这个指针来对数据库进行操作
var DB *sql.DB

// DBInit 是用来初始化
func DBInit(connString string) error {
	// connString是一个字符串，他保存着连接数据所需要的信息
	db, err := sql.Open("mysql", connString)
	// db 就是这个函数返回的一个连接mysql的一个实例
	if err != nil {
		panic(err)
	}
	DB = db

	return nil

}
