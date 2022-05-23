package forms

import "fmt"

func GetUserInput() (string, string, string, uint) {
	var (
		firstName, lastName, email string
		userTickets                uint
	)

	fmt.Printf("\nEnter your first name: ")
	fmt.Scanf("%v\n", &firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scanf("%v\n", &lastName)

	fmt.Printf("Enter your email: ")
	fmt.Scanf("%v\n", &email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scanf("%v\n", &userTickets)
	return firstName, lastName, email, userTickets
}
