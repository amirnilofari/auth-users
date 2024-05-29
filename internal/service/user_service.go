package service

import "github.com/amirnilofari/auth-users/internal/domain"

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) UserService {
	return UserService{repo: repo}
}

func (s UserService) CreateUser(firstName, lastName, email string) error {
	user := domain.User{Firstname: firstName, Lastname: lastName, Email: email}
	return s.repo.Save(user)
}

// Other business logic methods
