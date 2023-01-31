package layers

import "pustaka-api/models"

type Service interface {
	FindAll() ([]models.Book, error)
	FindByID(ID int) (models.Book, error)
	Create(book models.BookRequest) (models.Book, error)
	Update(ID int, book models.BookUpdateRequest) (models.Book, error)
	Delete(ID int) (models.Book, error)
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
	id, _ := bookRequest.ID.Int64()

	book := models.Book{
		ID:          int(id),
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest models.BookUpdateRequest) (models.Book, error) {
	book, err := s.repository.FindByID(ID)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	id, _ := bookRequest.ID.Int64()

	book.ID = int(id)
	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = int(price)
	book.Rating = int(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (models.Book, error) {
	book, err := s.repository.FindByID(ID)
	newBook, err := s.repository.Delete(book)
	return newBook, err
}
