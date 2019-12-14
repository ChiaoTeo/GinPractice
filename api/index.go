package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index 会返回主页
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "main website",
	})
}
