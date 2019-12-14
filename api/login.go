package api

import (
	"GinPractice/userinfo"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Login 是用来处理登陆操作的func
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" {
		fmt.Print("oooo")
	}
	fmt.Print(username)
	fmt.Print(password)
	if userinfo.CheckValid(username, password) {
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
