package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName   string
	lastName    string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	conferenceName := "GopherCon-2022"
	var remainingTickets uint = 100
	bookings := make([]UserData, 0)

	greetUser(conferenceName, remainingTickets)

	for {
		firstName, lastName, userTickets := getUserInput()
		isValidName, isValidUserTickets := helper.ValidateInput(firstName, lastName, userTickets, remainingTickets)
		if isValidName && isValidUserTickets {
			bookings, remainingTickets = bookTickets(bookings, firstName, lastName, userTickets, remainingTickets)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName)
		} else {

			continue
		}
		firstNames := getFirstNames(bookings)
		fmt.Println("The users who bought the tickets are: ")
		fmt.Println(firstNames)
		if remainingTickets == 0 {
			break
		}
	}
	wg.Wait()
}

func greetUser(conferenceName string, remainingTickets uint) {
	fmt.Println("Welcome to the", conferenceName, "Ticket Portal")
	fmt.Println("Book your tickets now, only", remainingTickets, "tickets left")
}

func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, user := range bookings {
		firstNames = append(firstNames, user.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter the Number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, userTickets
}

func bookTickets(bookings []UserData, firstName string, lastName string, userTickets uint, remainingTickets uint) ([]UserData, uint) {
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		userTickets: userTickets,
	}
	bookings = append(bookings, userData)
	remainingTickets -= userTickets
	fmt.Println("User ->", firstName+" "+lastName, ", booked", userTickets, "tickets")
	fmt.Println("Thankyou for booking the tickets,", remainingTickets, "left")
	return bookings, remainingTickets
}

func sendTicket(userTickets uint, firstName string, lastName string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("*************************************************************")
	fmt.Println("Sending the ticket:", ticket)
	fmt.Println("*************************************************************")
	wg.Done()
}
