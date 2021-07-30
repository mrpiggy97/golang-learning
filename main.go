package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generator(senderChannel chan<- int, waiter *sync.WaitGroup) {
	defer waiter.Done()
	time.Sleep(time.Second * 2)
	for i := 0; i <= 5; i++ {
		senderChannel <- rand.Int()
	}

	close(senderChannel)
	fmt.Println("generator is done")
}

func consumer(recieverChannel <-chan int, waiter *sync.WaitGroup, id int) {
	defer waiter.Done()
	time.Sleep(time.Second * 2)
	for task := range recieverChannel {
		var variableToPrint string = fmt.Sprintf("id %v task[%v]", id, task)
		fmt.Println(variableToPrint)
	}

	fmt.Println("consumer is done")
}

func worker(integerPointer *int) {
	for {
		time.Sleep(time.Millisecond * 500)
		*integerPointer = *integerPointer + 1
	}
}

func main() {

	rand.Seed(43)
	var channel chan int = make(chan int, 2)
	var waiter sync.WaitGroup

	waiter.Add(3)

	go generator(channel, &waiter)
	go consumer(channel, &waiter, 1)
	go consumer(channel, &waiter, 2)

	var firstTimer *time.Timer = time.NewTimer(5)
	var firstTicker *time.Ticker = time.NewTicker(time.Second)
	var initialValue int = 0
	var integerPointer *int = &initialValue

	waiter.Wait()

	go worker(integerPointer)

	for {
		select {
		case <-firstTimer.C:
			fmt.Printf("pointer value -> %d\n", *integerPointer)
		case <-firstTicker.C:
			fmt.Printf("pointer value -> %d\n", *integerPointer)
		}

		if *integerPointer >= 10 {
			break
		}
	}
}
