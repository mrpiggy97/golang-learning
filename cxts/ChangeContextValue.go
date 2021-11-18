package cxts

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

func ChangeContextValue(waiter *sync.WaitGroup, channel chan context.Context) {
	//create a new context with the values of the old context that will be delivered
	//by a channel
	defer waiter.Done()
	waiter.Add(1)
	var UUID ContextKey = "UUID"
	var UUIDValue float32 = rand.Float32()
	var message ContextKey = "message"
	println("it gets here")
	var oldContext context.Context = <-channel
	println("it gets gere")
	var newContext context.Context = context.WithValue(oldContext, UUID, UUIDValue)
	fmt.Printf("%v this is the old value\n", newContext.Value(message))
	fmt.Printf("this is the new value %v\n", newContext.Value(UUID))
	channel <- newContext
}
