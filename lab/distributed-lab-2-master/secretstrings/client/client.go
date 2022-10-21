package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	file, err := os.Open("wordlist")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		response := new(stubs.Response)
		request := stubs.Request{Message: scanner.Text()}
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Response:" + response.Message)
	}
	//request := stubs.Request{Message: "Hello"}

	//TODO: connect to the RPC server and send the request(s)
}
