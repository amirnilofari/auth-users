package cli

import (
	"fmt"

	"github.com/amirnilofari/auth-users/internal/service"
	"github.com/jedib0t/go-pretty/table"
)

func StartCLI(userService service.UserService) {
	// code to handle user input and call userService methods
	fmt.Println("User Managemet app")
	fmt.Println("1.Add a user")
	fmt.Println("2. See all users")
	fmt.Println("3. Search user by email")
	fmt.Println("4. Delete user")

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
	case selectedMainMenu == 2:
		users, err := userService.GetAllUsers()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			t := table.NewWriter()
			t.AppendHeader(table.Row{"#", "FirstName", "LastName", "Email"})
			for _, user := range users {
				t.AppendRow(table.Row{user.ID, user.Firstname, user.Lastname, user.Email})
			}
			fmt.Println(t.Render())
		}
	case selectedMainMenu == 3:
		var email string
		fmt.Print("Enter your desired user email: ")
		fmt.Scanln(&email)
		user, err := userService.GetUser(email)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("FirstName: ", user.Firstname)
			fmt.Println("LastName: ", user.Lastname)
		}
	case selectedMainMenu == 4:
		var id int
		fmt.Print("Enter your desired user id: ")
		fmt.Scanln(&id)
		err := userService.RemoveUser(id)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("The desired user was deleted!")
		}
	}

}
