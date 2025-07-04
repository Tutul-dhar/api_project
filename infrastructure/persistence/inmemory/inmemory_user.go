package inmemory

import (
	"errors"
	"github.com/tutul/book-server/domain/entity"
	"github.com/tutul/book-server/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type userRepo struct {
	store map[string]entity.User
}

func NewInMemoryUserRepo() repository.UserRepository {
	return &userRepo{store: make(map[string]entity.User)}
}

func (r *userRepo) Create(user entity.User) (entity.User, error) {
	if _, exists := r.store[user.Username]; exists {
		return entity.User{}, errors.New("user already exists")
	}
	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, err
	}
	user.Password = string(hashedPassword)
	r.store[user.Username] = user
	return user, nil
}

func (r *userRepo) GetByUsername(username string) (entity.User, error) {
	user, exists := r.store[username]
	if !exists {
		return entity.User{}, errors.New("user not found")
	}
	return user, nil
}
