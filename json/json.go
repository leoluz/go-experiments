package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Firstname string
	Lastname  string
}

func (person *Person) String() string {
	return fmt.Sprintf("%v %v\n", person.Firstname, person.Lastname)
}

func main() {
	leoluz := &Person{"Leonardo", "Luz"}
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
