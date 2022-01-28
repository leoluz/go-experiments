package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

func (person *Person) String() string {
	return fmt.Sprintf("%v %v\n", person.Firstname, person.Lastname)
}

func NewPerson(firstname, lastname string) (person *Person) {
	return &Person{firstname, lastname}
}

func NewCat(name string) *Person {
	return &Person{
		Firstname: name,
		Lastname:  "cat",
	}
}

func main() {
	leoluz := NewPerson("leo", "almeida")
	NewPerson("leo", "luz")
	fmt.Println(leoluz.String())
	personJson, err := json.MarshalIndent(leoluz, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(personJson))

	leoluzP := &Person{}
	err = json.Unmarshal(personJson, leoluzP)
	fmt.Println("Hi", leoluzP.Firstname)
}
