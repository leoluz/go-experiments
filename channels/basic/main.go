package main

import (
	"fmt"
)

func main() {

	messages := make(chan string)

	//messages <- "ping"
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
