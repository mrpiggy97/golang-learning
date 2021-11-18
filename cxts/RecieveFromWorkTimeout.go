package cxts

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func RecieveFromWorkTimeout(waiter *sync.WaitGroup, channel <-chan int, contextTimeout time.Duration) {
	defer waiter.Done()
	waiter.Add(1)
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout)

	select {
	case number := <-channel:
		fmt.Printf("time in worktimeout %v\n", number)
	case <-contextWithTimeout.Done():
		fmt.Println("context has reached timeout")
	}

	if contextWithTimeout.Err() != nil {
		fmt.Println(contextWithTimeout.Err())
	}

	cancel()
}
