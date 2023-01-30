package models

import (
	"encoding/json"
)

type BookEntity struct {
	Title string      `json:"title" binding:"required"`
	Years json.Number `json:"years" binding:"required,number"`
	Price json.Number `json:"price" binding:"required,number"`
}
