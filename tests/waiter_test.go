package tests

import (
	"testing"
	"time"
)

func senderFunc(channel chan<- int, number int) {
	channel <- number * 10
}

func SendNumber(channel chan int) {
	var recievedNumber int = <-channel
	channel <- recievedNumber * 10
	close(channel)
}

func TestWaiter(testCase *testing.T) {
	var channel chan int = make(chan int, 4)
	go senderFunc(channel, 10)
	go SendNumber(channel)
	time.Sleep(time.Second * 2)
	var number int = <-channel
	if number != 1000 {
		testCase.Error("number not 100")
	}
}
