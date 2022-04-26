package goroutines

import (
	"sync"
	"time"
)

func ChangeSharedString(mutexer *sync.Mutex, waiter *sync.WaitGroup, sharedString *string,
	slice *[]string, stringToAppend string) {
	//this function will write to a sharedString which is shared by multiple instances of
	//goroutines, ensuring each gets to change it
	defer waiter.Done()
	waiter.Add(1)
	for i := 0; i <= 3; i++ {
		mutexer.Lock()
		time.Sleep(time.Second * 1)
		*sharedString = stringToAppend
		var copySlice []string = *slice
		copySlice = append(copySlice, *sharedString)
		*slice = copySlice
		mutexer.Unlock()
	}
}
