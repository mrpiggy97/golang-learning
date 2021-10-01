package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func TickerReaction(tickerInstance *time.Ticker, waiter *sync.WaitGroup) {
	defer tickerInstance.Stop()
	defer waiter.Done()
	for {
		var currentTime time.Time = <-tickerInstance.C
		var currentSecond int = currentTime.Second()
		fmt.Printf("ticker time %v\n", currentSecond)
		if currentSecond >= 30 {
			return
		}
	}
}
