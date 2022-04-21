package mutexes

import "sync"

var balance int = 100

func Deposit(waiter *sync.WaitGroup, amount int, lock *sync.RWMutex) {
	lock.Lock()
	var currentBalance int = balance
	balance = currentBalance + amount
	lock.Unlock()
	defer waiter.Done()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	var currentBalance int = balance
	lock.RUnlock()
	return currentBalance
}
