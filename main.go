package main

import (
	"github.com/gin-gonic/gin"
	"pustaka-api/util"
)

func main() {
	router := gin.Default()

	router.GET("/", util.RootHandler)
	router.GET("/:id", util.IdHandler)
	router.GET("/book/query", util.QueryHandler) // http://localhost:8080/query?title=Halo
	router.POST("/book/post", util.PostBookHandler)

	router.Run()
}
