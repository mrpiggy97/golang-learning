package cxts

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func IncrementInteger(integerPointer *uint32, waiter *sync.WaitGroup) {
	defer waiter.Done()
	waiter.Add(1)
	var ticker *time.Ticker = time.NewTicker(time.Millisecond * 250)
	var deadline time.Time = time.Now().Add(time.Second * 3)
	deadlineContext, cancel := context.WithDeadline(context.Background(), deadline)
	for {
		select {
		case <-ticker.C:
			atomic.AddUint32(integerPointer, 1)
			fmt.Printf("%v this is the increment in uint32\n", *integerPointer)
		case <-deadlineContext.Done():
			fmt.Println("deadline context has reached its limit")
			fmt.Printf("%v\n", deadlineContext.Err().Error())
			ticker.Stop()
			cancel()
			return
		}
	}
}
