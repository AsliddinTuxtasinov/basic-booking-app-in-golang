package main

import (
	"fmt"

	"github.com/AsliddinTuxtasinon/booking-app/forms"
	"github.com/AsliddinTuxtasinon/booking-app/helper"
)

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = 50
	bookings              = make([]userDate, 0)
)

const conferenceTickets int = 50

type userDate struct {
	firstName, lastName, email string
	numberOfTickets            uint
}

func main() {
	greetUsers()

	for true {
		firstName, lastName, email, userTickets := forms.GetUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInputs(
			firstName, lastName, email, userTickets, remainingTickets,
		)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = bookTicket(firstName, lastName, email, userTickets)
			go helper.SendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("\nOur conferense is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			} else if !isValidEmail {
				fmt.Println("email adress you entered doesn't contain '@' or '.'")
			} else if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}

	}

}

func greetUsers() {
	fmt.Printf("\nWelcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available.\n\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets her to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(firstName, lastName, email string, userTickets uint) uint {
	remainingTickets -= userTickets

	var userDate = userDate{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userDate)

	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf(
		"\nThank you %v %v for booking %v tickets. You will recive a confirmation email at %v\n\n",
		firstName, lastName, userTickets, email,
	)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)

	return remainingTickets
}
