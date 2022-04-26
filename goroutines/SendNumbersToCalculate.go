package goroutines

import "time"

func SendNumbersToCalculate(channel chan<- int) {
	time.Sleep(time.Second * 3)
	channel <- 2222
	channel <- 211111
}
