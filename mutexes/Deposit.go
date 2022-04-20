package mutexes

import "sync"

var balance int = 100

func Deposit(waiter *sync.WaitGroup, amount int) {
	var currentBalance int = balance
	balance = currentBalance + amount
}

func Balance() int {
	var currentBalance int = balance
	return currentBalance
}
