package basicmodule

import (
	"errors"
	"strconv"
	"unicode/utf8"
)

const (
	_ticketCnt = 6
	_div       = 10
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
		cur := ticket % _div
		if i < _ticketCnt>>1 {
			lsum += cur
		} else {
			rsum += cur
		}
		ticket /= _div
	}
	return lsum == rsum, nil
}

// NHappyTickets returns the number of happy tickets in the range [a, b].
// return: count the number of happy tickets
func NHappyTickets(digits int) int64 {
	arr := sumArr(digits)
	if len(arr) == 0 {
		return 0
	}
	var rsl int64

	for _, v := range arr {
		rsl += v * v
	}
	return rsl
}

func sumArr(n int) []int64 {
	if n <= 0 {
		return nil
	}

	s := n*9 + 1
	arr := make([]int64, s)
	if n == 1 {
		for i := range arr {
			arr[i] = 1
		}
		return arr
	}
	prev := sumArr(n - 1)
	for i := range 10 {
		for j := range prev {
			if i+j < len(arr) {
				arr[i+j] += prev[j]
			}
		}
	}
	return arr
}

// Len func that returns the length of the string.
func Len(sym string) int {
	return utf8.RuneCountInString(sym)
}

func hasSixDigits(num int) bool {
	str := strconv.Itoa(num)
	return len(str) == _ticketCnt
}
