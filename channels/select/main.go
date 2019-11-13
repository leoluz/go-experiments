package main

import (
	"fmt"
	"time"
)

func main() {
	resultChan, doneChan := heavyProcess()

	done := false
	for !done {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timeout")
			done = true
		case result := <-resultChan:
			fmt.Println(result)
		case done = <-doneChan:
		}
	}
}

func heavyProcess() (<-chan string, <-chan bool) {
	resultChan := make(chan string)
	doneChan := make(chan bool)

	go do(resultChan, doneChan)
	return resultChan, doneChan
}

func do(result chan<- string, done chan<- bool) {
	for i := 0; i < 5; i++ {
		result <- fmt.Sprintf("Hi from %d", i)
		time.Sleep(time.Second * 1)
	}
	done <- true
	close(result)
	close(done)
}
