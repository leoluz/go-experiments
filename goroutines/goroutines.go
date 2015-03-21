package main

import (
	"fmt"
)

func main() {
	playPingPong()
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
