package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName string = "Go Conference"
	const allTickets = 50
	var remainingTickets uint = allTickets
	var bookings []string

	for remainingTickets <= allTickets {
		fmt.Printf("Welcome to our %v\n", conferenceName)
		fmt.Println("We have", remainingTickets, "tickets available")
		fmt.Println("Get your tickets here to attend")
		var userName string
		var lastName string
		var email string
		var userTickets int
		fmt.Println("What is your first name?")
		fmt.Scan(&userName)
		fmt.Println("What is your last name?")
		fmt.Scan(&lastName)
		fmt.Println("What is your email address?")
		fmt.Scan(&email)
		fmt.Println("How many tickets would you like to book?")
		fmt.Scan(&userTickets)

		isValidName := len(userName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, userName+" "+lastName)
			fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive confirmation on your email address - %v\n", userName, lastName, userTickets, email)
			fmt.Println("There are ", remainingTickets, "tickets available")

			firstNames := []string{}

			for _, booking := range bookings {
				names := strings.Fields(booking)
				fmt.Println(booking)
				firstNames = append(firstNames, names[0])
			}

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
