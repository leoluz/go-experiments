package main

import (
	"fmt"
	"time"
)

// main function
func main() {
	now := time.Now()
	//fmt.Printf("now: %v\n", now)
	t := now.UTC()
	fmt.Printf("t: %v\n", t)
	utc := now.In(time.UTC)
	fmt.Printf("utc: %v\n", utc)
}
