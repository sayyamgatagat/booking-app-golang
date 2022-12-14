package helper

func ValidateInput(firstName string, lastName string, userTickets uint, remainingTickets uint) (bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidUserTickets := remainingTickets >= userTickets
	return isValidName, isValidUserTickets
}
