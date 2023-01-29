package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "Mirza Hilmi",
			"Bio": "A Software Developer, Backend Developing and Robotics Enthusiast",
		})
	})  // <- Router Path
}