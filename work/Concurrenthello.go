package main

import (
	"fmt"
	_ "fmt"
	"time"
	_ "time"
)

func Hello() {

	for i := 0; i <= 4; i++ {
		go fmt.Println("Hello from goroutine:", i)
	}
}

func main() {
	Hello()
	time.Sleep(1 * time.Second)
}
