package api

import (
	"GinPractice/users"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Login 是用来处理登陆操作的func
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" {
		fmt.Print("后台消息“不存在的用户要登陆")
	}
	fmt.Print(username)
	fmt.Print(password)
	success, err := users.CheckVaild(username, password)
	if err != nil {
		c.JSON(200, gin.H{
			"status":  "10002",
			"message": "数据库错误，请联系管理人员",
		})

	} else {
		if success {
			c.JSON(200, gin.H{
				//0代表无错误
				"status":  "0",
				"message": "信息正确，成功登录",
			})
		} else {
			c.JSON(200, gin.H{
				"status":  "00001",
				"message": "用户名或密码不正确",
			})
		}
	}

}
