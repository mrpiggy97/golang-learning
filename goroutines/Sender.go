package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func Sender(senderChannel chan<- string, waiter *sync.WaitGroup) {
	defer waiter.Done()
	defer close(senderChannel)
	time.Sleep(time.Second * 3)
	var stringToSend string = "this is the message of sender"
	senderChannel <- stringToSend
	fmt.Println("sender goroutine finished")
}
