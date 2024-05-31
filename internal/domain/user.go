package domain

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
}

type UserRepository interface {
	Save(user User) error
	FindByEmail(Email string) (User, error)
	FindAll() ([]User, error)
	//DeletByID(id int) error
}
