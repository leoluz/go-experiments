package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	list, err := ioutil.ReadDir("/Users/leo.almeida/dev/go/src/github.com/leoluz")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	for _, file := range list {
		fmt.Printf("%s (is dir:%v)\n", file.Name(), file.IsDir())
	}
}
