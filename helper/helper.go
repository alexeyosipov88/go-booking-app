package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateUserInput(userName string, lastName string, email string, userTickets int, remainingTickets uint) (bool1 bool, bool2 bool, bool3 bool) {

	isValidName := len(userName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}

func BookTicket(remainingTickets uint, userName string, lastName string, userTickets int, email string, bookings []map[string]string) []map[string]string {

	remainingTickets = remainingTickets - uint(userTickets)

	var userData = make(map[string]string)

	fullName := userName + " " + lastName
	userData["fullName"] = fullName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	fmt.Printf("Thank you, %v %v for booking %v tickets. You will receive confirmation on your email address - %v\n", userName, lastName, userTickets, email)
	fmt.Println("There are ", remainingTickets, "tickets available")
	bookings = append(bookings, userData)

	return bookings

}
