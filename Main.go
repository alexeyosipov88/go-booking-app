package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const allTickets = 50
	var remainingTickets uint = allTickets
	var bookings []string

	for {
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

		remainingTickets = remainingTickets - uint(userTickets)

		fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive confirmation on your email address - %v\n", userName, lastName, userTickets, email)
		fmt.Println("There are ", remainingTickets, "tickets available")

		bookings = append(bookings, userName+" "+lastName)
		fmt.Printf("All bookings: %v\n", bookings)

	}

}
