package main

import "fmt"

type A struct {
	Name     string
	LastName string
}

type B struct {
	Name string
	age  int
}

type Person struct {
	A
	B
}

func main() {
	p := &Person{}
	p.Name = "Leo"
	fmt.Printf("p.A.Name=%v\np.B.Name=%v", p.A.Name, p.B.Name)
}
