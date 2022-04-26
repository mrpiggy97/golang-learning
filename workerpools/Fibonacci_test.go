package workerpools_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-learning/workerpools"
)

func TestFibonacci(testCase *testing.T) {
	var tasks []int = []int{2, 4, 6, 7, 8, 22}
	var maxWorkerNumber int = 3
	var jobsChannel chan int = make(chan int, 2)
	var resultsChannel chan int = make(chan int, 2)
	for i := 0; i < maxWorkerNumber; i++ {
		go workerpools.FiboWorker(i, jobsChannel, resultsChannel)
	}
	for _, task := range tasks {
		jobsChannel <- task
	}
	close(jobsChannel)
	for {
		result, channelISAvailable := <-resultsChannel
		if channelISAvailable {
			fmt.Println(result)
			if result >= 17711 {
				close(resultsChannel)
			}
		} else {
			break
		}
	}
}
