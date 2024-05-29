package repository

import (
	"database/sql"
	"errors"

	"github.com/amirnilofari/auth-users/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Save(user domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (first_name, last_name, email) VALUES (?,?)", user.Firstname, user.Lastname, user.Email)
	return err
}

func (r *UserRepository) FindByID(id int) (domain.User, error) {
	var user domain.User
	row := r.db.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found!")
		}
		return user, err
	}
	return user, nil
}

// other methods to implement domain.UserRepository
