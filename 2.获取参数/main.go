package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		username := c.DefaultQuery("usrname", "lcq")
		pwd := c.Query("pwd")
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"username": username,
			"pwd":      pwd,
		})
	})
	r.Run()
}
