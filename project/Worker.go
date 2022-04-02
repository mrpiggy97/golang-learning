package project

import (
	"fmt"
	"time"

	"github.com/mrpiggy97/golang-learning/workerpools"
)

type Worker struct {
	Id          int
	JobQueue    chan Job
	WorkerPool  chan chan Job
	QuitChannel chan bool
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:          id,
		WorkerPool:  workerPool,
		JobQueue:    make(chan Job, 1),
		QuitChannel: make(chan bool, 1),
	}
}

// Start will run a goroutine that will run an infinite for
// loop, inside that loop we will first send values to
// workerIntance.WorkerPool, then we will wait for value to come
// through workerInstance.JobQueue channel
func (workerInstance *Worker) Start() {
	go func() {
		for {
			workerInstance.WorkerPool <- workerInstance.JobQueue
			select {
			case job := <-workerInstance.JobQueue:
				fmt.Printf("worker with id %v\n", workerInstance.Id)
				fib := workerpools.Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("worker with id %v finished with result %v\n", workerInstance.Id, fib)
			case <-workerInstance.QuitChannel:
				fmt.Printf("worker with id %v has finished\n", workerInstance.Id)
			}
		}
	}()
}

func (workerIns *Worker) Stop() {
	go func() {
		workerIns.QuitChannel <- true
	}()
}
