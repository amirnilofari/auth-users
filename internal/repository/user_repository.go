package repository

import (
	"errors"

	"github.com/amirnilofari/auth-users/internal/domain"
)

type UserRepository struct {
	users map[int]domain.User
}

func NewUserRepository() domain.UserRepository {
	return &UserRepository{users: make(map[int]domain.User)}
}

func (r *UserRepository) Save(user domain.User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}

// other methods to implement domain.UserRepository
