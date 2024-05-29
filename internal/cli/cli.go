package cli

import (
	"fmt"

	"github.com/amirnilofari/auth-users/internal/service"
)

func StartCLI(userService service.UserService) {
	// code to handle user input and call userService methods
	fmt.Println("User Managemet app")
	fmt.Println("1.Add a user")
	fmt.Println("2. See all users")

	var selectedMainMenu int
	fmt.Print("Choose number of options: ")
	fmt.Scanln(&selectedMainMenu)

	switch {
	case selectedMainMenu == 1:
		fmt.Println("Enter the desired user information:")
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

}

// func mainMenu(userService service.UserService) {

// }
