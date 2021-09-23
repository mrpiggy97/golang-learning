package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func generator(senderChannel chan<- int, waiter *sync.WaitGroup) {
	defer waiter.Done()
	defer close(senderChannel)
	time.Sleep(time.Second * 2)
	for i := 0; i <= 5; i++ {
		senderChannel <- rand.Int()
	}
	fmt.Println("generator is done")
}

func consumer(recieverChannel <-chan int, waiter *sync.WaitGroup, id int) {
	defer waiter.Done()
	time.Sleep(time.Second * 4)
	for task := range recieverChannel {
		var variableToPrint string = fmt.Sprintf("id %v task[%v]", id, task)
		fmt.Println(variableToPrint)
	}

	fmt.Println("consumer is done")
}

func tickerReaction(tickerInstance *time.Ticker, waiter *sync.WaitGroup) {
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

func timerReaction(timerInstance *time.Timer, waiter *sync.WaitGroup) {
	defer timerInstance.Stop()
	defer waiter.Done()
	var currentTime time.Time = <-timerInstance.C
	var currentSecond int = currentTime.Second()
	fmt.Printf("timer time %v\n", currentSecond)
}

func reciever(recieverChannel <-chan string, waiter *sync.WaitGroup) {
	defer waiter.Done()
	time.Sleep(time.Second * 1)
	var byteChain string = <-recieverChannel
	fmt.Println(byteChain)
}

func sender(senderChannel chan<- string, waiter *sync.WaitGroup) {
	defer waiter.Done()
	defer close(senderChannel)
	time.Sleep(time.Second * 3)
	var stringToSend string = "this is the message of sender"
	senderChannel <- stringToSend
	fmt.Println("sender goroutine finished")
}

func setter(integerPointer *int32, contextInstance context.Context, id int, waiter *sync.WaitGroup) {
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

func triggerCancel(cancelF context.CancelFunc, waiter *sync.WaitGroup) {
	defer waiter.Done()
	time.Sleep(time.Second * 1)
	fmt.Println("trigger cancel trigger cancel trigger cancel")
	cancelF()
}

func workTimeout(sleepTime time.Duration, info chan<- int, waiter *sync.WaitGroup) {
	defer waiter.Done()
	waiter.Add(1)
	time.Sleep(sleepTime)
	info <- int(sleepTime)
}

func recieveFromWorkTimeout(waiter *sync.WaitGroup, channel <-chan int, contextTimeout time.Duration) {
	defer waiter.Done()
	waiter.Add(1)
	contextWithTimeout, cancel := context.WithTimeout(context.Background(), contextTimeout)

	select {
	case number := <-channel:
		fmt.Printf("time in worktimeout %v\n", number)
	case <-contextWithTimeout.Done():
		fmt.Println("context has reached timeout")
	}

	if contextWithTimeout.Err() != nil {
		fmt.Println(contextWithTimeout.Err())
	}

	cancel()
}

func incrementInteger(integerPointer *uint32, waiter *sync.WaitGroup) {
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

type messageKey string
type UUIDKey string
type operationKey string

func makeContextValue(waiter *sync.WaitGroup, channel chan<- context.Context) {
	defer waiter.Done()
	defer close(channel)
	waiter.Add(1)
	time.Sleep(time.Second * 3)
	var message messageKey = "message"
	var messageValue string = "this is the message"
	newContext := context.WithValue(context.Background(), message, messageValue)
	fmt.Printf("%v\n", newContext.Value(message))
	channel <- newContext
}
func changeContextValue(waiter *sync.WaitGroup, channel <-chan context.Context) {
	defer waiter.Done()
	waiter.Add(1)
	var UUID UUIDKey = "UUID"
	var message messageKey = "message"
	var UUIDValue float32 = rand.Float32()
	var oldContext context.Context = <-channel
	var newContext context.Context = context.WithValue(oldContext, UUID, UUIDValue)
	fmt.Printf("%v this is the old value\n", newContext.Value(message))
	fmt.Printf("this is the new value %v\n", newContext.Value(UUID))
}

func sendNumbersToCalculate(channel chan<- int) {
	time.Sleep(time.Second * 3)
	channel <- 2222
	channel <- 211111
}

func doCalculation(number1 int, number2 int, waiter *sync.WaitGroup, channel <-chan int) {
	defer waiter.Done()
	waiter.Add(1)
	var operation operationKey = "operation"
	cxt := context.WithValue(context.Background(), operation, "+")
	cxt2, cancel := context.WithDeadline(cxt, time.Now().Add(time.Millisecond*100))
	time.Sleep(time.Second * 2)
	var initialValue int = 0
	var total *int = &initialValue

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

var first int = 0

func RunOnce(channel chan<- bool, number int, onceInstance *sync.Once, waiter *sync.WaitGroup) {
	defer waiter.Done()
	waiter.Add(1)
	var randomInt uint32 = rand.Uint32() % 100
	fmt.Printf("%v this is the random number\n", randomInt)
	time.Sleep(time.Duration(randomInt) * time.Millisecond)
	onceInstance.Do(func() {
		first = number
		channel <- true
		close(channel)
	})
	fmt.Println("run once done")
}

func main() {
	var timerInstance *time.Timer = time.NewTimer(time.Second * 3)
	var tickerInstance *time.Ticker = time.NewTicker(time.Second)
	var waiter *sync.WaitGroup = &sync.WaitGroup{}
	var channel chan int = make(chan int, 2)
	var channel2 chan string = make(chan string)
	var integer int32 = 0
	var workTimeoutChannel chan int = make(chan int, 2)
	var contextTimeout time.Duration = time.Millisecond * 800
	var initialValue uint32 = 2330
	var deadlineIntegerPointer *uint32 = &initialValue

	var contextValueChannel chan context.Context = make(chan context.Context, 2)

	contextCancel, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go setter(&integer, contextCancel, i, waiter)
	}
	for sleepTime := 250; sleepTime <= 1000; sleepTime = sleepTime + 250 {
		var sleepTime time.Duration = time.Duration(sleepTime * int(time.Millisecond))
		go workTimeout(sleepTime, workTimeoutChannel, waiter)
		go recieveFromWorkTimeout(waiter, workTimeoutChannel, contextTimeout)
	}
	go incrementInteger(deadlineIntegerPointer, waiter)
	go makeContextValue(waiter, contextValueChannel)
	go changeContextValue(waiter, contextValueChannel)
	var channelCalc chan int = make(chan int, 2)
	go sendNumbersToCalculate(channelCalc)
	go doCalculation(4000, 5000, waiter, channelCalc)
	var onceChannel chan bool = make(chan bool, 2)
	var onceInstance *sync.Once = &sync.Once{}
	for i := 1; i <= 5; i++ {
		go RunOnce(onceChannel, first, onceInstance, waiter)
	}
	for valueInChannel := range onceChannel {
		fmt.Printf("%v\n", valueInChannel)
	}
	waiter.Add(7)
	go timerReaction(timerInstance, waiter)
	go tickerReaction(tickerInstance, waiter)
	go generator(channel, waiter)
	go consumer(channel, waiter, 2)
	go reciever(channel2, waiter)
	go sender(channel2, waiter)
	go triggerCancel(cancel, waiter)
	waiter.Wait()
	fmt.Printf("final count %v\n", integer)
	fmt.Println("main program finished")
}
