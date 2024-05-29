package cli

import (
	"fmt"
	"github.com/amirnilofari/auth-users/internal/service"
)

func StartCLI(userService service.UserService) {
	// code to handle user input and call userService methods
	fmt.Println("Welcome to the user management CLI")
	var firstName, lastName, email string
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)
	err := userService.CreateUser(firstName, lastName, email)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User created successfully!")
	}
}
