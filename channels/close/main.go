package main

import "fmt"

func main() {
	nChan := getNumbers()

	//for {
	//number, isOpen := <-nChan
	//if isOpen {
	//fmt.Println(number)
	//} else {
	//break
	//}
	//}

	for number := range nChan {
		fmt.Println(number)
	}
}

func getNumbers() <-chan int {
	nChan := make(chan int)
	go produce(nChan)
	return nChan
}

func produce(nChan chan<- int) {
	nChan <- 1
	nChan <- 2
	nChan <- 3
	close(nChan)
}
