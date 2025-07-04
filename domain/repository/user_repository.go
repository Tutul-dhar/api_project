package repository

import "github.com/tutul/book-server/domain/entity"

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	GetByUsername(username string) (entity.User, error)
}
