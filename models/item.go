package models

type BookEntity struct {
	Title string      `json:"title" binding:"required"`
	Years interface{} `json:"years" binding:"required,number"`
	Price interface{} `json:"price" binding:"required,number"`
}
