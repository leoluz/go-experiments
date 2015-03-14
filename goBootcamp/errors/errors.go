package main

import (
	"fmt"
)

func main() {
	typeAssertion()
}

func typeAssertion() {
	printHeader("Type Assertion")
	var name interface{}
	name = "Leo"
	if me, ok := name.(string); !ok {
		fmt.Println("Var name is not a int")
	} else {
		fmt.Println("Var name is a string:", me)
	}
}

func printHeader(title string) {
	fmt.Println("------------------------")
	fmt.Println(title, ":\n")
}
