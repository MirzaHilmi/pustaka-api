package main

import (
	"fmt"
	"log"
	"pustaka-api/layers"
	"pustaka-api/util"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=@y0N90D1NG dbname=pustaka port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	} else {
		fmt.Println("DB connection successfull")
	}

	bookRepo := layers.NewRepo(db)
	bookService := layers.NewService(bookRepo)
	bookHandler := util.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("v1")

	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/:id", bookHandler.IdHandler)
	v1.GET("/book/query", bookHandler.QueryHandler) // http://localhost:8080/query?title=Halo
	v1.POST("/book/post", bookHandler.PostBookHandler)
	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.PUT("/books/edit/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/delete/:id", bookHandler.DeleteBook)

	router.Run()
}

//main
//handler
//service
//repository
//database
