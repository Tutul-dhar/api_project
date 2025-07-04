package service

import (
	"errors"
	"github.com/tutul/book-server/domain/entity"
	"github.com/tutul/book-server/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Register(user entity.User) (entity.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) Login(username, password string) (entity.User, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return entity.User{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return entity.User{}, errors.New("invalid password")
	}
	return user, nil
}
