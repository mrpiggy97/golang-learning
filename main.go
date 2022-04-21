package main

import (
	"fmt"
	"sync"

	"github.com/mrpiggy97/golang-learning/mutexes"
)

// Main will apply all concepts learned.

func main() {
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	var mutexer *sync.RWMutex = new(sync.RWMutex)
	for i := 1; i <= 5; i++ {
		waiter.Add(1)
		go mutexes.Deposit(waiter, i*100, mutexer)
	}
	waiter.Wait()
	fmt.Println(mutexes.Balance(mutexer))
}
