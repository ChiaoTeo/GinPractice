package conf

import (
	"GinPractice/model"

	_ "github.com/go-sql-driver/mysql"
)

// DBInit 是用来初始化
func DBInit(connString string) error {
	err := model.Init(connString)
	if err != nil {
		return err
	}
	return nil
}

// DBClose 关闭数据库链接
func DBClose() {
	model.DB.Close()
}
