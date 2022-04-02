package project

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorkers: maxWorkers,
		WorkerPool: make(chan chan Job, maxWorkers),
	}
}

// Dispatch will run an infinite for loop, inside that loop
// we will consume dispatchInstance.JobQueue channel,
// if channel is available we run goroutine that will,
// consume dispatchInstance.WorkerPool through a variable
// declaration and then we will send data through that channel
func (dispatchInstance *Dispatcher) Dispatch() {
	for {
		select {
		case job, channelAvailable := <-dispatchInstance.JobQueue:
			if !channelAvailable {
				break
			}
			go func() {
				workerJobQueue := <-dispatchInstance.WorkerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (dispatcherInstance *Dispatcher) Run() {
	for i := 0; i <= dispatcherInstance.MaxWorkers; i++ {
		var newWorker *Worker = NewWorker(i, dispatcherInstance.WorkerPool)
		newWorker.Start()
	}
	go dispatcherInstance.Dispatch()
}
