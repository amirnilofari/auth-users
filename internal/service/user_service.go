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

func (s UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s UserService) GetUser(email string) (domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s UserService) RemoveUser(id int) error {
	return s.repo.DeletByID(id)
}
