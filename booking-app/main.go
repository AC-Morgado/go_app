package main

import (
	"booking-app/userInfo"
	"fmt"
	"math/rand"
	"time"
)

// Variable declarations
var conferenceName = "Go Conference"
const CONFERENCE_TICKETS uint = 50
var remainingTickets uint = CONFERENCE_TICKETS
var bookings []UserData

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

type Ticket struct {
	id         uint
	firstName  string
	lastName   string
	seatNumber uint
	time       time.Time
}

func main() {
	
	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets uint
	
	greetUsers()

	for remainingTickets > 0 {
		userFirstName, userLastName, userEmail, userTickets = userInfo.GetUserInput(remainingTickets)
		fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s shortly.\n", userFirstName, userLastName, userTickets, userEmail)
		bookTicket(userFirstName, userLastName, userEmail, userTickets)
	}

	fmt.Println("Our conference is fully booked. Come back next year!")
	printBookings()

}

func bookTicket(firstName, lastName, email string, tickets uint) {

	// Update remaining tickets
	remainingTickets = remainingTickets - tickets
	
	// Store booking
	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   tickets,
	}
	bookings = append(bookings, userData)

	for range tickets {

		// Create a new ticket
		var ticket = Ticket{
			id:         uint(rand.Uint32()),
			firstName:  firstName,
			lastName:   lastName,
			seatNumber: uint(len(bookings) + 1),
			time:       time.Now(),
		}
		sendTicket(ticket)
	}

	fmt.Println("---------------------------------------------------")
	fmt.Printf("All tickets have been sent to %s!\n", email)
	fmt.Println("---------------------------------------------------")	

}

func greetUsers() {
	fmt.Printf("Welcome to our %s booking application! :)\n", conferenceName)
	fmt.Printf("The total number of tickets is %d, and there are %d tickets remaining.\n", CONFERENCE_TICKETS, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printBookings() {
	userFirstNames := []string{}
	for _, booking := range bookings {
		var firstName string
		fmt.Sscanf(booking.firstName, "%s", &firstName)
		userFirstNames = append(userFirstNames, firstName)
	}
}

func sendTicket(ticket Ticket) {
	time.Sleep(6 * time.Second)
}
