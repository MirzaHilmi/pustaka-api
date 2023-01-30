package main

import (
	"github.com/gin-gonic/gin"
	"pustaka-api/util"
)

func main() {
	router := gin.Default()

	v1 := router.Group("v1")

	v1.GET("/", util.RootHandler)
	v1.GET("/:id", util.IdHandler)
	v1.GET("/book/query", util.QueryHandler) // http://localhost:8080/query?title=Halo
	v1.POST("/book/post", util.PostBookHandler)

	router.Run()
}
