package main

import (
	"flag"
	"sync"

	"github.com/mrpiggy97/golang-learning/net"
)

// Main will apply all concepts learned.

func main() {
	flag.Parse()
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	for i := 0; i <= 100; i++ {
		waiter.Add(1)
		go net.StartPortScanning(waiter, i)
	}
	waiter.Wait()
}
