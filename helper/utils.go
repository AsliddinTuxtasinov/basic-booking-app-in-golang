package helper

import (
	"fmt"
	"strings"
	"time"
)

// Validate User Inputs
func ValidateUserInputs(firstName, lastName, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func SendTicket(userTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Printf("\n####################\n")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Printf("\n####################\n")
}
