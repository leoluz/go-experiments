package main

import "fmt"

func main() {
	nChan := make(chan int, 3)
	getNumbers(nChan)

	fmt.Println(<-nChan, <-nChan, <-nChan)
}

func getNumbers(nChan chan int) {
	nChan <- 1
	nChan <- 2
	nChan <- 3
}
