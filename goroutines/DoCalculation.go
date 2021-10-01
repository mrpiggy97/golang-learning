package goroutines

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/mrpiggy97/golang-learning/cxts"
)

func DoCalculation(number1 int, number2 int, waiter *sync.WaitGroup, channel <-chan int) {
	defer waiter.Done()
	waiter.Add(1)
	var operation cxts.ContextKey = "operation"
	cxt := context.WithValue(context.Background(), operation, "+")
	cxt2, cancel := context.WithDeadline(cxt, time.Now().Add(time.Millisecond*100))
	time.Sleep(time.Second * 2)
	var total *int = new(int)

	fmt.Printf("%v done shit\n", <-cxt2.Done())

	select {
	case <-cxt2.Done():
		fmt.Println("do calculation log deadline has reached its limit")
	case x := <-channel:
		y := <-channel
		*total = x + y
		fmt.Printf("do calculation log total is %v\n", *total)
	}

	cancel()
}
