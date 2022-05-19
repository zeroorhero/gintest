package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/book/:username/:pwd", func(c *gin.Context) {
		username := c.Param("username")
		pwd := c.Param("pwd")
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"pwd":      pwd,
		})
	})
	router.Run()
}
