package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	group := router.Group("/book")
	{
		group.GET("/bug", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ok",
			})
		})
		group.GET("/xxx", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"xxx": "ok",
			})
		})
	}
	router.Run()
}
