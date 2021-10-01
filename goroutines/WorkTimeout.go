package goroutines

import (
	"sync"
	"time"
)

func WorkTimeout(sleepTime time.Duration, info chan<- int, waiter *sync.WaitGroup) {
	defer waiter.Done()
	waiter.Add(1)
	time.Sleep(sleepTime)
	info <- int(sleepTime)
}
