package main

import (
	"fmt"
	"log"
	"pustaka-api/models"
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
	}
	fmt.Println("DB connection successfull")

	book := models.Book{}
	// book.ID = 3
	// book.Title = "Gamata"
	// book.Description = "Gam"
	// book.Price = 358900
	// book.Rating = 5

	// err = db.Debug().Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error creating book record")
	// }
	// fmt.Println("Book have been successfuly recorded")

	err = db.Where("id = ?", 3).First(&book).Error

	if err != nil {
		fmt.Println("Error creating book record")
	}

	book.Description = "Wind, cold yet so calming"
	err = db.Save(&book).Error

	if err != nil {
		fmt.Println("Error editing book record")
	}

	fmt.Println("Book have been successfuly edited")

	router := gin.Default()

	v1 := router.Group("v1")

	v1.GET("/", util.RootHandler)
	v1.GET("/:id", util.IdHandler)
	v1.GET("/book/query", util.QueryHandler) // http://localhost:8080/query?title=Halo
	v1.POST("/book/post", util.PostBookHandler)

	router.Run()
}
