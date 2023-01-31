package layers

import "pustaka-api/models"

type Service interface {
	FindAll() ([]models.Book, error)
	FindByID(ID int) (models.Book, error)
	Create(book models.BookRequest) (models.Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]models.Book, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (models.Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(bookRequest models.BookRequest) (models.Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := models.Book{
		Title: bookRequest.Title,
		Description: bookRequest.Description,
		Price: int(price),
		Rating: int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}
