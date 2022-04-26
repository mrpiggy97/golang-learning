package main

import (
	"sync"

	"github.com/mrpiggy97/golang-learning/workerpools"
)

// Main will apply all concepts learned.

func main() {
	service := workerpools.NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))
	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}
