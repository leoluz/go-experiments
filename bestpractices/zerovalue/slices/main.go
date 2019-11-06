package main

import "fmt"

type Person struct {
	Name string
}

type SliceTest struct {
	PersonByValues   []Person
	PersonByPointers []*Person
}

func main() {

	p1 := Person{Name: "person 1"}

	sliceTest := SliceTest{
		PersonByValues:   []Person{p1, p1},
		PersonByPointers: []*Person{&p1, &p1},
	}
	sliceTest.PersonByValues[0].Name = "person value modified"
	sliceTest.PersonByPointers[0].Name = "person pointer modified"

	fmt.Printf("person: %s\n", p1.Name)
	print(sliceTest)
}

func print(sliceTest SliceTest) {
	fmt.Printf("Person.Value[0]: %s\n", sliceTest.PersonByValues[0].Name)
	fmt.Printf("Person.Value[1]: %s\n", sliceTest.PersonByValues[1].Name)
	fmt.Printf("Person.Pointer[0]: %s\n", sliceTest.PersonByPointers[0].Name)
	fmt.Printf("Person.Pointer[1]: %s\n", sliceTest.PersonByPointers[1].Name)
}
