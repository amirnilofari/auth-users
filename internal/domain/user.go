package domain

type User struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
}

type UserRepository interface {
	Save(user User) error
	FindByID(id int) (User, error)
	//FindAll() ([]User, error)
	//DeletByID(id int) error
}
