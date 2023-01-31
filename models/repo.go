package models

import "gorm.io/gorm"

type Repository interface {
	FindAll([]Book, error)
	FindByID(id int) (Book, error)
	CreateBook(book Book) (Book, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repo) FindByID(ID int) (Book, error) {
	var book Book
	err := r.db.First(&book, ID).Error
	return book, err
}

func (r *repo) CreateBook(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}
