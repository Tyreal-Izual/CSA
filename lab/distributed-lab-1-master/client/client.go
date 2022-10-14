package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
	}
	//TODO In a continuous loop, read a message from the server and display it.
}

func write(conn net.Conn) {
	writer := bufio.NewWriter(conn)
	fmt.Printf("Enter text->")
	msg, _ := writer.WriteString()
	fmt.Println(conn, msg)

	//TODO Continually get input from the user and send messages to the server.
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	for {
		conn, _ := net.Dial("tcp", "127.0.0.1:8030")
		go read(conn)
		go write(conn)
		fmt.Fprintf(conn, *addrPtr)
		read(conn)
	}
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
