package net

import (
	"flag"
	"fmt"
	"net"
	"sync"

	"github.com/rs/zerolog/log"
)

var siteFlag *string = flag.String("site", "", "sets site to test scanning port")

func StartPortScanning(waiter *sync.WaitGroup, port int) {
	defer waiter.Done()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *siteFlag, port))
	if err != nil {
		return
	}
	var closingErr error = conn.Close()
	if closingErr != nil {
		log.Error().Msg(closingErr.Error())
		return
	}
	fmt.Printf("port %d is open\n", port)
}
