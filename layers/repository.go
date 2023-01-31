package layers

import (
	"gorm.io/gorm"
	"pustaka-api/models"
)

type Repository interface {
	FindAll() ([]models.Book, error)
	FindByID(ID int) (models.Book, error)
	Create(book models.Book) (models.Book, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repo) FindByID(ID int) (models.Book, error) {
	var book models.Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *repo) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}
