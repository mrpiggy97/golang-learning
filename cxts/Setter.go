package cxts

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Setter(integerPointer *int32, contextInstance context.Context, id int, waiter *sync.WaitGroup) {
	defer waiter.Done()
	waiter.Add(1)
	var tickerInstance *time.Ticker = time.NewTicker(time.Millisecond * 250)
	time.Sleep(time.Millisecond * 100)

	for {
		select {
		case <-contextInstance.Done():
			fmt.Printf("%v done done done done\n", id)
			tickerInstance.Stop()
			return

		case <-tickerInstance.C:
			fmt.Println("ticker instance is ready")
			fmt.Printf("%v\n", *integerPointer)
			atomic.AddInt32(integerPointer, 1)
		}
	}
}
