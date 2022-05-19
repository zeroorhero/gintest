package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	Name string `form:"user"`
	Pwd  int    `form:"pass"`
}

func main() {
	router := gin.Default()
	router.GET("/book", func(c *gin.Context) {
		var b book
		if err := c.ShouldBind(&b); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message":  "ok",
				"username": b.Name,
				"pwd":      b.Pwd,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	})
	router.Run()
}
