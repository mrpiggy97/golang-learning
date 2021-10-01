package goroutines

import (
	"sync"
)

func IncreaseValue(value *int, waiterInstance *sync.WaitGroup, mutexInstance *sync.Mutex) {
	defer waiterInstance.Done()
	waiterInstance.Add(1)
	mutexInstance.Lock()
	*value += 1
	mutexInstance.Unlock()
}

func RecieveFromChannel(waiterInstance *sync.WaitGroup, channel <-chan int, value *int) {
	defer waiterInstance.Done()
	waiterInstance.Add(1)
	for {
		newVal, channelAvailable := <-channel
		if channelAvailable {
			*value = *value + newVal
		} else {
			break
		}
	}
}

func SendNewValue(waiterInstance *sync.WaitGroup, channel chan<- int, value *int) {
	defer waiterInstance.Done()
	waiterInstance.Add(1)
	for i := 0; i <= 99; i++ {
		channel <- 1
	}
	close(channel)
}
