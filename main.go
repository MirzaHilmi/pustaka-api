package main

import (
	"fmt"
	"log"
	"pustaka-api/util"
	"pustaka-api/layers"

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

	books, err := bookRepo.FindAll()

	if err != nil {
		log.Fatal("Error Occured")
	} else {
		fmt.Println("Successfull")
	}

	for _, book := range books {
		fmt.Printf("ID: %v\nTitle: %v\v", book.ID, book.Title)
	} 

	router := gin.Default()

	v1 := router.Group("v1")

	v1.GET("/", util.RootHandler)
	v1.GET("/:id", util.IdHandler)
	v1.GET("/book/query", util.QueryHandler) // http://localhost:8080/query?title=Halo
	v1.POST("/book/post", util.PostBookHandler)

	router.Run()
}
