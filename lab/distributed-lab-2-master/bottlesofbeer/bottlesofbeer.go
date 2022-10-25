package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"time"
	// "net/rpc"
	// "fmt"
	// "time"
	// "net"
)

var nextAddr string
var nextround *rpc.Client
var initialized = false

func Beer(i int) {
	time.Sleep(1 * time.Second)
	if i > 1 {
		fmt.Printf("%v bottles of beer on the wall, %v bottles of beer. Take one down, pass it around...", i, i)
	} else if i == 1 {
		fmt.Printf("1 bottles of beer on the wall, 1 bottles of beer. Take one down, pass it around...")
	} else {
		fmt.Printf("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall...")
	}

}

func PassItAround(bottles int) {
	request := Token{Bottles: bottles}
	response := new(Token)
	if initialized == false {
		nextround, _ = rpc.Dial("tcp", nextAddr)
		initialized = true
	}
	nextround.Go("BottlesOfBeer.Round", request, response, nil)
}

type BottlesOfBeer struct{}
type Token struct {
	Bottles int
}

func (b *BottlesOfBeer) Round(intoken Token, outtoken *Token) (err error) {
	bottles := intoken.Bottles
	Beer(bottles)
	if bottles > 0 {
		PassItAround(bottles - 1)
	}
	return
}
func main() {
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	bottles := flag.Int("n", 0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	rpc.Register(&BottlesOfBeer{})
	listener, _ := net.Listen("tcp", ":"+*thisPort)
	defer listener.Close()
	if *bottles > 0 {
		Beer(*bottles)
		go PassItAround(*bottles - 1)
	}
	rpc.Accept(listener)
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.
}
