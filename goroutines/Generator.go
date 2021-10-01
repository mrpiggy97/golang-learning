package goroutines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Generator(senderChannel chan<- int, waiter *sync.WaitGroup) {
	defer waiter.Done()
	defer close(senderChannel)
	time.Sleep(time.Second * 2)
	for i := 0; i <= 5; i++ {
		senderChannel <- rand.Int()
	}
	fmt.Println("generator is done")
}
