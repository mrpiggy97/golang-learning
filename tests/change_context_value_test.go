package tests

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/mrpiggy97/golang-learning/cxts"
)

func testReadPreviousContextValue(testCase *testing.T) {
	//check that values of previous context are still present in the new
	//context
	var channel chan context.Context = make(chan context.Context, 1)
	var messageKey cxts.ContextKey = "message"
	var messageVal string = "this belongs to the previous context"
	var previousContext context.Context = context.WithValue(context.Background(),
		messageKey, messageVal)
	channel <- previousContext
	var preWaiter *sync.WaitGroup = new(sync.WaitGroup)
	go cxts.ChangeContextValue(preWaiter, channel)
	time.Sleep(time.Millisecond * 500)
	var nextContext context.Context = <-channel
	if nextContext.Value(messageKey) != messageVal {
		var errMessage string = fmt.Sprintf("nextContext.value(messageKey) should return %v", messageVal)
		testCase.Error(errMessage)
	}

}

func testReadNewContextValue(testCase *testing.T) {
	//test that a new context is indeed sent through channel
	//and that we can access its values
	var channel chan context.Context = make(chan context.Context, 1)
	var messageKey cxts.ContextKey = "message"
	var messageVal string = "this belongs to the previous context"
	var UUIDKey cxts.ContextKey = "UUID"
	var previousContext context.Context = context.WithValue(context.Background(),
		messageKey, messageVal)
	channel <- previousContext
	var testWaiter *sync.WaitGroup = new(sync.WaitGroup)
	go cxts.ChangeContextValue(testWaiter, channel)
	//we sleep to give time for goroutine to access the channel,
	//otherwise we are stuck waiting because the following line
	//after sleep consumes the value that is supposed to be
	//for the goroutine in testing
	time.Sleep(time.Millisecond * 500)
	var newContext context.Context = <-channel
	if newContext.Value(UUIDKey) == nil {
		fmt.Println(newContext.Value(UUIDKey))
		testCase.Error("UUID should not be nil")
	}
}

func TestChangeContextValue(testCase *testing.T) {
	testCase.Run("Action=read-previous-context", testReadPreviousContextValue)
	testCase.Run("Action=read-next-context", testReadNewContextValue)
}
