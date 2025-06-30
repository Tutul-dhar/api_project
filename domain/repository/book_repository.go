package repository

import "book-server/domain/entity"

type BookRepository interface {
	Create(book entity.Book) (entity.Book, error)
	GetByID(id string) (entity.Book, error)
	List() ([]entity.Book, error)
	Update(id string, book entity.Book) (entity.Book, error)
	Delete(id string) (entity.Book, error)
}
