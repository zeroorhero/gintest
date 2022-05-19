package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/book", func(c *gin.Context) {
		username := c.PostForm("username")
		pwd := c.DefaultPostForm("pwd", "222")
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"username": username,
			"pwd":      pwd,
		})
	})
	r.Run()
}
