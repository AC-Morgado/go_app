package userInfo

import (
	"fmt"
	"strings"
)

func GetUserInput(remainingTickets uint) (string, string, string, uint) {

	fmt.Println("Please enter your first name: ")
	userFirstName := getValidName()

	fmt.Println("Please enter your last name: ")
	userLastName := getValidName()

	fmt.Println("Please enter your email address: ")
	userEmail := getValidEmail()

	fmt.Println("Please enter the number of tickets you want to book: ")
	userTickets := getValidTicketNumber(remainingTickets)

	return userFirstName, userLastName, userEmail, userTickets
}

func getValidName() string {
	var name string
	fmt.Scan(&name)

	//Check first name not null and doesn't have numbers or special characters
	var isValidFirstName bool = false
	for !isValidFirstName {
		if len(name) < 2 {
			fmt.Print("First name must be at least 2 characters long. Please enter your first name: ")
			fmt.Scan(&name)
		} else if !isValidName(name) {
			fmt.Print("First name must not contain numbers or special characters. Please enter your first name: ")
			fmt.Scan(&name)
		} else {
			isValidFirstName = true
			break
		}
	}
	return name
}

func getValidEmail() string {
	var userEmail string
	fmt.Scan(&userEmail)

	//Check for valid email
	var isValidEmail bool = false
	for !isValidEmail {
		if len(userEmail) < 7 {
			fmt.Print("Email address must be at least 7 characters long. Please enter your email address: ")
			fmt.Scan(&userEmail)
		} else if !strings.Contains(userEmail, "@") {
			fmt.Print("Email address must contain '@' symbol. Please enter your email address: ")
			fmt.Scan(&userEmail)
		} else if !strings.Contains(userEmail, ".") {
			fmt.Print("Email address must contain '.' symbol. Please enter your email address: ")
			fmt.Scan(&userEmail)
		} else {
			isValidEmail = true
			break
		}
	}
	return userEmail
}

func getValidTicketNumber(remainingTickets uint) uint {
	var userTickets uint
	fmt.Scan(&userTickets)

	//Check ticket availability
	var isValidTicketNumber bool = false
	for !isValidTicketNumber {
		if userTickets <= 0 {
			fmt.Print("Number of tickets must be at least 1. Please enter a valid number of tickets: ")
			fmt.Scan(&userTickets)
		} else if userTickets > remainingTickets {
			fmt.Printf("We only have %d tickets remaining, so you can't book %d tickets.\n", remainingTickets, userTickets)
			fmt.Print("Please enter a valid number of tickets: ")
			fmt.Scan(&userTickets)
		} else {
			isValidTicketNumber = true
			break
		}
	}
	return userTickets
}

func isValidName(name string) bool {
	for _, char := range name {
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			return false
		}
	}
	return true
}
