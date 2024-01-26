package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{"title": "Main Page"},
		)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	r.Run(":8090")
}
