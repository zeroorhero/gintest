package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"lcq": "zrf",
		})
	})
	g.Run()
}
