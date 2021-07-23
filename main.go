package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func helloB(channel chan<- string) {
	fmt.Println("helloB started execution")
	time.Sleep(time.Second * 1)
	channel <- "this is mister b"
	channel <- "this is mister b 2.0"
	channel <- "this is mister b 3.0"
}

func Reciever(channel <-chan string) {
	fmt.Println("Reciever started execution")
	time.Sleep(time.Second * 2)
	for message := range channel {
		fmt.Println(message)
	}
}

func sendNumbers(num chan<- int) {
	fmt.Println("sendNumbers started execution")
	for i := 0; i <= 2; i++ {
		time.Sleep(time.Second * 1)
		num <- i
		if i == 2 {
			close(num)
		}
	}
}

func sendString(str chan<- string) {
	fmt.Println("sendString started execution")
	for i := 0; i <= 3; i++ {
		time.Sleep(time.Second * 1)
		var iConverted string = fmt.Sprintf("this is the number %v", i)
		str <- iConverted
		if i == 3 {
			close(str)
		}
	}
}

func recieveNumber(num <-chan int) {
	fmt.Println("recieveNumber started execution")
	time.Sleep(time.Second * 5)
	for {
		number, channelIsOpen := <-num
		if channelIsOpen {
			fmt.Println(number)
		} else {
			fmt.Println("no more numbers")
			break
		}
	}
}

func recieveString(str <-chan string) {
	fmt.Println("recieveString started execution")
	time.Sleep(time.Second * 6)
	for {
		byteChain, channelIsOpen := <-str
		if channelIsOpen {
			fmt.Println(byteChain)
		} else {
			fmt.Println("no more strings")
			break
		}
	}
}

func sum(num1 uint32, num2 uint32, channel chan<- uint32) {
	fmt.Println("sum started execution")
	time.Sleep(time.Second * 8)
	channel <- num1 + num2
}

func substract(num1 uint32, num2 uint32, channel chan<- uint32) {
	fmt.Println("substract started execution")
	time.Sleep(time.Second * 7)
	channel <- num1 - num2
}

func checkRuntimeExecution(channel <-chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println(<-channel)
}

func generator(channel chan<- int, waiter *sync.WaitGroup) {
	defer waiter.Done()
	time.Sleep(time.Second * 2)
	for i := 0; i <= 5; i++ {
		channel <- rand.Int()
	}

	close(channel)
	fmt.Println("generator is done")
}

func consumer(channel <-chan int, waiter *sync.WaitGroup, id int) {
	defer waiter.Done()
	time.Sleep(time.Second * 2)
	for task := range channel {
		var variableToPrint string = fmt.Sprintf("id %v task[%v]", id, task)
		fmt.Println(variableToPrint)
	}
}

func main() {

	var firstChannel chan string = make(chan string, 2)
	var secondChannel chan int = make(chan int, 2)
	var thirdChannel chan string = make(chan string, 2)
	var fourthChannel chan uint32 = make(chan uint32, 2)
	var fifthChannel chan uint32 = make(chan uint32, 2)
	var sixthChannel chan string = make(chan string, 2)

	sixthChannel <- "this is the value of the sixth channel"

	go helloB(firstChannel)
	go Reciever(firstChannel)
	go sendNumbers(secondChannel)
	go recieveNumber(secondChannel)
	go sendString(thirdChannel)
	go recieveString(thirdChannel)

	go sum(100, 100, fourthChannel)
	go substract(200, 100, fifthChannel)
	go checkRuntimeExecution(sixthChannel)

	select {
	case result := <-fourthChannel:
		fmt.Println("addition has been successful")
		fmt.Println(result)
	case result := <-fifthChannel:
		fmt.Println("substraction has been successful")
		fmt.Println(result)
	default:
		fmt.Println("no gorutine has finished")
	}

	time.Sleep(time.Second * 10)
}
