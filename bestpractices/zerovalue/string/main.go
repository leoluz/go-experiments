package main

import (
	"fmt"
)

func main() {

	//var name *string
	name := new(string)

	if *name == "" {
		fmt.Println("string zero value")
	} else {
		fmt.Println("not initialized")
	}
}
