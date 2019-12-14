package main

import (
	"GinPractice/api"

	"github.com/gin-gonic/gin"
)

// "GinPractice/api"
// "github.com/gin-gonic/gin

func main() {
	// router 是一个gin.Engine类型的指针，也就是说这个Default()做了这样一件事：
	//初始化了一个gin的引擎，然后把这个引擎的指针返回，这样我们可以通过router对这个引擎进行操作
	router := gin.Default()
	//这个地方是说：如果要找网页（我们做好的html）应该取哪里找，在这里我把网页放到了web文件夹中
	router.LoadHTMLGlob("web/*")
	//这个句话是说，如果直接访问"localhost:8080"会做什么事情
	router.GET("/", api.Index)
	//router.POST(arg1,arg2) 做了这样一个操作，通过router这个指针告诉gin.Engine，
	//如果有人来发一个get请求，我将会进行的一个操作，这个操作就是这个GET函数的第二次参数arg2
	//我们注意到第二个参数是一个函数，是一个接受一个gin.Context类型，并且返回值为空的函数
	//也就是：我们在设置路由的时候会用到GET(),POST(),PUT()...等函数，
	//第一个参数是用string,代表请求的路由，第二个是一个接受一个gin.Context类型，并且返回值为空的函数
	router.POST("/user/login", api.Login)

	router.Run(":8080")

}
