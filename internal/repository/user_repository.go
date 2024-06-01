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
	_, err := r.db.Exec("INSERT INTO users (first_name, last_name, email) VALUES (?,?,?)", user.Firstname, user.Lastname, user.Email)
	return err
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (r *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	row := r.db.QueryRow("SELECT * FROM users WHERE email = ?", email)
	if err := row.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("user not found!")
		}
		return user, err
	}
	return user, nil
}

func (r *UserRepository) DeletByID(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
