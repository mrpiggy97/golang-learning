package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RunOnce(channel chan<- bool, number int, onceInstance *sync.Once, waiter *sync.WaitGroup,
	first int) {
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
