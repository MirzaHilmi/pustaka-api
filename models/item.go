package models

type Book struct {
	Title string `json:"title"`
	Years uint   `json:"years"`
	Price uint   `json:"price"`
}
