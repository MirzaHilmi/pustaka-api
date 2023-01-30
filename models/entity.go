package models

import "time"

type BookTest struct {
	Title string      `json:"title" binding:"required"`
	Years interface{} `json:"years" binding:"required,number"`
	Price interface{} `json:"price" binding:"required,number"`
}

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
