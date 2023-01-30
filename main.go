package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/:id", idHandler)
	router.GET("/query", proxyHandler) // http://localhost:8080/query?title=Halo
	router.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Mirza Hilmi",
		"Bio":  "A Software Developer, Backend Developing and Robotics Enthusiast",
	})
}

func idHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func proxyHandler(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{"title": title})
}
