package inmemory

import (
	"book-server/domain/entity"
	"book-server/domain/repository"
	"errors"
)

type bookRepo struct {
	store map[string]entity.Book
}

func NewInMemoryRepo() repository.BookRepository {
	return &bookRepo{store: make(map[string]entity.Book)}
}

func (r *bookRepo) Create(book entity.Book) (entity.Book, error) {
	r.store[book.UUID] = book
	return book, nil
}

func (r *bookRepo) GetByID(id string) (entity.Book, error) {
	book, exists := r.store[id]
	if !exists {
		return entity.Book{}, errors.New("not found")
	}
	return book, nil
}

func (r *bookRepo) List() ([]entity.Book, error) {
	var books []entity.Book
	for _, book := range r.store {
		books = append(books, book)
	}
	return books, nil
}

func (r *bookRepo) Update(id string, book entity.Book) (entity.Book, error) {
	if _, exists := r.store[id]; !exists {
		return entity.Book{}, errors.New("not found")
	}
	r.store[id] = book
	return book, nil
}

func (r *bookRepo) Delete(id string) (entity.Book, error) {
	book, exists := r.store[id]
	if !exists {
		return entity.Book{}, errors.New("not found")
	}
	delete(r.store, id)
	return book, nil
}
