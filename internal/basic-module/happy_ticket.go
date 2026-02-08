package basicmodule

import (
	"errors"
	"strconv"
)

const (
	_ticketCnt = 6
	div        = 10
)

var errTicketLimitDigits = errors.New("ticket must have 6 digits")

// IsHappyTicket checks if the ticket number is happy.
// params: ticket - the ticket number to check.
// return: bool, error - if true then the ticket is happy, if false then the ticket is not happy.
func IsHappyTicket(ticket int) (bool, error) {
	if !hasSixDigits(ticket) {
		return false, errTicketLimitDigits
	}
	lsum := 0
	rsum := 0
	for i := range _ticketCnt {
		cur := ticket % div
		if i < _ticketCnt>>1 {
			lsum += cur
		} else {
			rsum += cur
		}
		ticket /= div
	}
	return lsum == rsum, nil
}

func hasSixDigits(num int) bool {
	str := strconv.Itoa(num)
	return len(str) == _ticketCnt
}
