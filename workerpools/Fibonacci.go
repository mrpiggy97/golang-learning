package workerpools

import (
	"fmt"
	"sync"
	"time"
)

func FiboWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker with id %v started fib %v\n", id, job)
		var res int = Fibonacci(job)
		results <- res
	}
}

func Fibonacci(n int) int {
	time.Sleep(time.Second * 2)
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Locker     *sync.RWMutex
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
		Locker:     new(sync.RWMutex),
	}
}

func (serviceInstance *Service) Work(job int) {
	serviceInstance.Locker.RLock()
	exists := serviceInstance.InProgress[job]
	if exists {
		serviceInstance.Locker.RUnlock()
		var response chan int = make(chan int, 1)
		defer close(response)
		serviceInstance.Locker.Lock()
		serviceInstance.IsPending[job] = append(serviceInstance.IsPending[job], response)
		serviceInstance.Locker.Unlock()
		fmt.Printf("waiting response for job %v\n", job)
		var res int = <-response
		fmt.Printf("response recieved %d\n", res)
		return
	}
	serviceInstance.Locker.RUnlock()
	serviceInstance.Locker.Lock()
	serviceInstance.InProgress[job] = true
	serviceInstance.Locker.Unlock()
	fmt.Printf("Calculate Fibonacci for %d\n", job)
	result := Fibonacci(job)
	serviceInstance.Locker.RLock()
	pendingWorkers, exists := serviceInstance.IsPending[job]
	serviceInstance.Locker.RUnlock()
	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result
		}
		fmt.Printf("Result sent - all pending workers ready job:%d\n", job)
	}
	serviceInstance.Locker.Lock()
	serviceInstance.InProgress[job] = false
	serviceInstance.IsPending[job] = make([]chan int, 0)
	serviceInstance.Locker.Unlock()
}
