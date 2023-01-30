package models

type Book struct {
	Title string `json:"title" binding:"required"`
	Years uint   `json:"years" binding:"required, number"`
	Price uint   `json:"price" binding:"required, number"`
}
