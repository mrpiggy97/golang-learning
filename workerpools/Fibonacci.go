package workerpools

import "fmt"

func FiboWorker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker with id %v started fib %v\n", id, job)
		var res int = Fibonacci(job)
		results <- res
		if job >= 22 {
			close(results)
		}
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
