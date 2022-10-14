package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8083")
	if err != nil {
		//handle error
	}
	fmt.Println(conn, "err")
	status, err := bufio.NewReader(conn).ReadString('\n')
	ln, err := net.Listen("tcp", "127.0.0.1:8083")
	for {
		conn, err := ln.Accept()
		if err != nil {
			//handle error
		}
	}
	//TODO: all
	//Deal with an error event.
}

func acceptConns(ln net.Listener, conns chan net.Conn) {

	reader := bufio.NewReader(conns)
	conns <- reader()

	ln, _ = net.Listen("tcp", "127.0.0.1:8030")
	for {
		conns, _ = ln.Accept()
		conns <- handleClient()
	}
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	reader := bufio.NewReader(msgs)

	msg, _ := reader.ReadString('\n')

	fmt.Println(msg)
	fmt.Fprintln(client, "OK")
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()
	Listener := bufio.NewListener(os.Stdin)

	//TODO Create a Listener for TCP connections on the port given above.

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
		}
	}
}
