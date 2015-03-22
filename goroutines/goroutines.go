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
	var worker sync.WaitGroup
	worker.Add(3)
	go printSeconds(5, &worker)
	go printSeconds(10, &worker)
	go printSeconds(3, &worker)
	worker.Wait()
}

func printSeconds(sec int, worker *sync.WaitGroup) {
	defer worker.Done()
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
