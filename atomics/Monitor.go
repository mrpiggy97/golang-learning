package atomics

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Monitor struct {
	ActiveUsers int
	Requests    int
}

func Update(checker atomic.Value, mutexer *sync.Mutex, waiterInstance *sync.WaitGroup) {
	defer waiterInstance.Done()
	waiterInstance.Add(1)
	for i := 0; i <= 9; i++ {
		time.Sleep(time.Millisecond * 500)
		mutexer.Lock()
		current := checker.Load().(*Monitor)
		current.ActiveUsers += 100
		current.Requests += 300
		checker.Store(current)
		mutexer.Unlock()
	}
}

func Observe(checker atomic.Value, waiterInstance *sync.WaitGroup) {
	defer waiterInstance.Done()
	waiterInstance.Add(1)
	for i := 0; i <= 9; i++ {
		time.Sleep(time.Second * 1)
		current := checker.Load().(*Monitor)
		fmt.Printf("this is the monitor %v\n", current)
	}
}
