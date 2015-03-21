package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// shows how to use select with channels
	playPingPong()

	// shows how to sync multiple goroutines waiting untill all are done
	syncGroup()
}

func syncGroup() {
	var control sync.WaitGroup
	control.Add(3)
	go printSeconds(5, &control)
	go printSeconds(10, &control)
	go printSeconds(3, &control)
	control.Wait()
}

func printSeconds(sec int, control *sync.WaitGroup) {
	defer control.Done()
	for i := 0; i <= sec; i++ {
		fmt.Print(i, " ")
		time.Sleep(time.Second)
	}
}

func ping(c chan<- string, done chan<- bool) {
	for i := 0; i < 5; i++ {
		c <- "ping"
	}
	done <- true
}

func pong(c chan<- string, done chan<- bool) {
	for i := 0; i < 5; i++ {
		c <- "pong"
	}
	done <- true
}

func playPingPong() {
	c1, c2 := make(chan string), make(chan string)
	cDonePing, cDonePong := make(chan bool), make(chan bool)

	go ping(c1, cDonePing)
	go pong(c2, cDonePong)

	donePing, donePong := false, false
	for !donePing || !donePong {
		select {
		case c := <-c1:
			fmt.Println("Print from ping:", c)
		case c := <-c2:
			fmt.Println("Print from pong:", c)
		case donePing = <-cDonePing:
			fmt.Println("Done with Ping")
		case donePong = <-cDonePong:
			fmt.Println("Done with Pong")
		}
	}
}
