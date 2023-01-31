package models

import (
	"encoding/json"
	"time"
)

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookRequest struct {
	ID          json.Number `json:"id" binding:"required"`
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"required,number"`
}

type BookUpdateRequest struct {
	ID          json.Number `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Price       json.Number `json:"price"`
	Rating      json.Number `json:"rating"`
}

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}
