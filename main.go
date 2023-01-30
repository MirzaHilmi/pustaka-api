package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"pustaka-api/models"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/:id", idHandler)
	router.GET("/book/query", queryHandler) // http://localhost:8080/query?title=Halo
	router.POST("/book/post", postBookHandler)
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

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func postBookHandler(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		log.Fatal("Error occured, please try again")
	}

	c.JSON(http.StatusOK, book)
}
