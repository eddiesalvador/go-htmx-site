package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create new Gin router
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Define a route
	router.GET("/oi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Oi, lua!!",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H {
			"title": "Main website",
		})
	})

	router.Run("localhost:8080")
}
