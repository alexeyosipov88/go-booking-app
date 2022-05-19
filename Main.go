package main

import "fmt"

func main() {

	var name string = "Alexey"
	const conferenceTickets = 50
	var conferenceName string = "Go Conference"
	remainingTickets := 0

	fmt.Printf("Welcome to our %v, %v\n", conferenceName, name)
	fmt.Println("We have", remainingTickets, "tickets available")

	fmt.Println("Get your tickets here to attend")
	fmt.Println(conferenceTickets)

}
