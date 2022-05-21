package main

import (
	"fmt"
	"go-booking-app/helper"
	"strings"
)

var conferenceName string = "Go Conference"

const allTickets = 50

var bookings = make([]map[string]string, 0)

var remainingTickets uint = allTickets

func main() {

	greetUsers()
	for remainingTickets <= allTickets {
		var userName string
		var lastName string
		var email string
		var userTickets int

		userName, lastName, email, userTickets = getUserInput(userName, lastName, email, userTickets)
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userName, lastName, email, allTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookings = helper.BookTicket(remainingTickets, userName, lastName, userTickets, email, bookings)

			firstNames := getFirstNames(bookings)

			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			var noTicketsRemaining bool = remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Sorry, the name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Sorry, an email you entered is not valid")
			}
			if !isValidTicketNumber {
				fmt.Println("Sorry, the number of tickets you entered is invalid")
			}

		}
	}

}

func greetUsers() {
	fmt.Println("Welcome to our conference:", conferenceName)
	fmt.Println("We have", remainingTickets, "tickets available")

}

func getFirstNames(bookings []map[string]string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking["fullName"])
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput(userName string, lastName string, email string, userTickets int) (string, string, string, int) {

	fmt.Println("What is your first name?")
	fmt.Scan(&userName)
	fmt.Println("What is your last name?")
	fmt.Scan(&lastName)
	fmt.Println("What is your email address?")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&userTickets)

	return userName, lastName, email, userTickets

}
