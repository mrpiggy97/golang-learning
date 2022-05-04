package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
)

type Client chan<- string

var (
	incomingClients chan Client = make(chan Client)
	leavingClients  chan Client = make(chan Client)
	messages        chan string = make(chan string)
)

var (
	host *string = flag.String("host", "localhost", "sets host")
	port *int    = flag.Int("port", 3090, "sets port")
)

func HandleConnection(conn net.Conn) {
	var message chan string = make(chan string)
	defer conn.Close()
	go MessageWrite(conn, message)
	var clientName string = conn.RemoteAddr().String()
	message <- fmt.Sprintf("welcome to the server %s", clientName)
	messages <- fmt.Sprintf("client is here, name %s", clientName)
	incomingClients <- message
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s:%s\n", clientName, inputMessage.Text())
	}
	leavingClients <- message
	messages <- fmt.Sprintf("%s said goodbye!", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}

func BroadCast(waiter *sync.WaitGroup) {
	defer waiter.Done()
	var clients map[Client]bool = make(map[Client]bool)
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalf(err.Error())
	}
	var waiter *sync.WaitGroup = new(sync.WaitGroup)
	waiter.Add(1)
	go BroadCast(waiter)
	for {
		conn, connErr := listener.Accept()
		if connErr != nil {
			log.Print(connErr.Error())
			continue
		}
		waiter.Add(1)
		go HandleConnection(conn)
	}
}
