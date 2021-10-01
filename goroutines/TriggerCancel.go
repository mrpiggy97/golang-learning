package goroutines

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func TriggerCancel(cancelF context.CancelFunc, waiter *sync.WaitGroup) {
	defer waiter.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("trigger cancel trigger cancel trigger cancel")
	cancelF()
}
