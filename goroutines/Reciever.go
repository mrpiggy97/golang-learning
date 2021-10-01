package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Reciever(recieverChannel <-chan string, waiter *sync.WaitGroup) {
	defer waiter.Done()
	time.Sleep(time.Second * 1)
	var byteChain string = <-recieverChannel
	fmt.Println(byteChain)
}
