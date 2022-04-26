package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func TimerReaction(timerInstance *time.Timer, waiter *sync.WaitGroup) {
	defer timerInstance.Stop()
	defer waiter.Done()
	var currentTime time.Time = <-timerInstance.C
	var currentSecond int = currentTime.Second()
	fmt.Printf("timer time %v\n", currentSecond)
}
