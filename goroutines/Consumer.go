package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Consumer(recieverChannel <-chan int, waiter *sync.WaitGroup, id int) {
	defer waiter.Done()
	time.Sleep(time.Second * 4)
	for task := range recieverChannel {
		var variableToPrint string = fmt.Sprintf("id %v task[%v]", id, task)
		fmt.Println(variableToPrint)
	}

	fmt.Println("consumer is done")
}
