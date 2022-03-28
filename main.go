package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mrpiggy97/golang-learning/ErrorHandling"
	"github.com/mrpiggy97/golang-learning/InOut"
	"github.com/mrpiggy97/golang-learning/atomics"
	"github.com/mrpiggy97/golang-learning/cxts"
	"github.com/mrpiggy97/golang-learning/embedding"
	"github.com/mrpiggy97/golang-learning/encodings"
	"github.com/mrpiggy97/golang-learning/goroutines"
	"github.com/mrpiggy97/golang-learning/stringManipulation"
)

// Main will apply all concepts learned.

func main() {
	ErrorHandling.Something()
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
		go cxts.Setter(&integer, contextCancel, i, waiter)
	}
	for sleepTime := 250; sleepTime <= 1000; sleepTime = sleepTime + 250 {
		var sleepTime time.Duration = time.Duration(sleepTime * int(time.Millisecond))
		go goroutines.WorkTimeout(sleepTime, workTimeoutChannel, waiter)
		cxts.RecieveFromWorkTimeout(waiter, workTimeoutChannel, contextTimeout)
	}
	go cxts.IncrementInteger(deadlineIntegerPointer, waiter)
	go cxts.MakeContextValue(waiter, contextValueChannel)
	go cxts.ChangeContextValue(waiter, contextValueChannel)
	var channelCalc chan int = make(chan int, 2)
	go goroutines.SendNumbersToCalculate(channelCalc)
	go cxts.DoCalculation(4000, 5000, waiter, channelCalc)
	var onceChannel chan bool = make(chan bool, 2)
	var onceInstance *sync.Once = &sync.Once{}
	var first int = 0
	for i := 1; i <= 5; i++ {
		go goroutines.RunOnce(onceChannel, i, onceInstance, waiter, first)
	}
	for valueInChannel := range onceChannel {
		fmt.Printf("%v\n", valueInChannel)
	}

	var stringToBeShared *string = new(string)
	var stringSlice *[]string = new([]string)
	var mutexInstance *sync.Mutex = new(sync.Mutex)
	go goroutines.ChangeSharedString(mutexInstance, waiter, stringToBeShared, stringSlice, "fabian")
	go goroutines.ChangeSharedString(mutexInstance, waiter, stringToBeShared, stringSlice, "chris")
	waiter.Add(7)
	go goroutines.TimerReaction(timerInstance, waiter)
	go goroutines.TickerReaction(tickerInstance, waiter)
	go goroutines.Generator(channel, waiter)
	go goroutines.Consumer(channel, waiter, 2)
	go goroutines.Reciever(channel2, waiter)
	go goroutines.Sender(channel2, waiter)
	go cxts.TriggerCancel(cancel, waiter)
	var val *int = new(int)
	for i := 0; i <= 100; i++ {
		go goroutines.IncreaseValue(val, waiter, mutexInstance)
	}
	var value2 *int = new(int)
	var channelLock chan int = make(chan int, 101)
	go goroutines.SendNewValue(waiter, channelLock, value2)
	go goroutines.RecieveFromChannel(waiter, channelLock, value2)

	var intCounter *int32 = new(int32)
	go atomics.Increaser(intCounter)
	go atomics.Decreaser(intCounter)
	for i := 0; i <= 5; i++ {
		fmt.Printf("%v this is the loaded value\n", atomic.LoadInt32(intCounter))
		time.Sleep(time.Millisecond * 500)
	}
	var checkerValue atomic.Value = atomic.Value{}
	var monitorInstance atomics.Monitor = atomics.Monitor{0, 0}
	checkerValue.Store(&monitorInstance)
	go atomics.Update(checkerValue, mutexInstance, waiter)
	go atomics.Observe(checkerValue, waiter)

	var target []byte = make([]byte, 30)
	empty := InOut.Reader{}
	numb, err := empty.Read(target)

	mr := &InOut.Reader{Data: "save the world with go", From: 2}
	n, er := mr.Read(target)

	for i := 0; i < len(mr.Data); i++ {
		b := byte(mr.Data[i]) == mr.Data[i]
		fmt.Printf("%v value of character in string %v\n", byte(mr.Data[i]), b)
	}

	var message []byte = []byte("this is a byte string")
	fmt.Printf("%v this is the byte slice\n", message)
	var MyWriter InOut.Writer = InOut.Writer{"save ", 6}
	si, e := MyWriter.Write(message)
	var stringSlicer []string = make([]string, 100)
	for i := 0; i < len(stringSlicer); i++ {
		stringSlicer[i] = fmt.Sprintf("%v %v", "fabian", i)
	}
	InOut.WriteToFile(stringSlicer, "/tmp/numbers.txt")
	InOut.WriteToFile(stringSlicer, "/tmp/numbers.txt")
	InOut.ReadFile("/tmp/numbers.txt")
	InOut.OsWriter("/workspaces/golang-learning/InOut/write-here.txt", stringSlicer[0:10])
	InOut.ReadFile("/tmp/numbers.txt")
	InOut.OsOffSet("/tmp/numbers.txt")
	waiter.Wait()
	fmt.Printf("final count %v\n", integer)
	fmt.Printf("%v\n", *stringSlice)
	fmt.Printf("%v this is the new value\n", *val)
	fmt.Printf("%v\n", *value2)
	fmt.Printf("%v\n", atomic.LoadInt32(intCounter))
	fmt.Printf("%v bytes %v error\n", numb, err)
	fmt.Printf("%v bytes %v error %v target\n", n, er, target)
	fmt.Printf("%v %v\n", MyWriter.Data, si)
	fmt.Printf("%v\n", e)
	var stringToBeManipulated string = "faBiAnJeEsus"
	stringManipulation.StringToUpperCase(stringToBeManipulated)
	InOut.WriterStd()
	InOut.ReaderStd()
	InOut.BufferReader()
	InOut.Scanner()
	InOut.BufferWriter()
	InOut.UserInput()
	InOut.BufferUserInput()
	encodings.CSVExample()
	encodings.CSVWrite()
	encodings.AsOriginal()
	encodings.PrintJSON()
	fmt.Println("main program finished")
	var student1 embedding.Student = embedding.Student{
		Person: embedding.Person{Id: 11222, Name: "fabian"},
		Class:  "4b",
	}
	student1.PrintName()
}
