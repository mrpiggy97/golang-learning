package cxts

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func MakeContextValue(waiter *sync.WaitGroup, channel chan<- context.Context) {
	defer waiter.Done()
	waiter.Add(1)
	time.Sleep(time.Second * 3)
	var message ContextKey = "message"
	var messageValue string = "this is the message"
	newContext := context.WithValue(context.Background(), message, messageValue)
	fmt.Printf("%v\n", newContext.Value(message))
	channel <- newContext
}
