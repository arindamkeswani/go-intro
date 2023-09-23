package helper

import (
    "strings"
)

func ValidateUserInput(
	firstName string,
	lastName string,
	email string,
	userTickets uint,
	remainingTickets uint,
) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNo := userTickets > 0 && userTickets <= remainingTickets

	// isValidUser := isValidName && isValidEmail && isValidTicketNo;

	return isValidName, isValidEmail, isValidTicketNo

}

