package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

var (
	clientPort *int    = flag.Int("port", 3090, "sets port to be used")
	clientHost *string = flag.String("host", "localhost", "sets host")
)

func main() {
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", *clientHost, *clientPort))
	if err != nil {
		log.Fatalf(err.Error())
	}
	done := make(chan struct{})
	waiter.Add(1)
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
		defer waiter.Done()
	}()
	CopyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
