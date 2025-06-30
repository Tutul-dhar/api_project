package service

import (
	"book-server/domain/entity"
	"book-server/domain/repository"
	"github.com/google/uuid"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{repo}
}

func (s *BookService) Create(book entity.Book) (entity.Book, error) {
	book.UUID = uuid.NewString()
	return s.repo.Create(book)
}

func (s *BookService) Get(id string) (entity.Book, error) {
	return s.repo.GetByID(id)
}

func (s *BookService) List() ([]entity.Book, error) {
	return s.repo.List()
}

func (s *BookService) Update(id string, book entity.Book) (entity.Book, error) {
	book.UUID = id
	return s.repo.Update(id, book)
}

func (s *BookService) Delete(id string) (entity.Book, error) {
	return s.repo.Delete(id)
}
