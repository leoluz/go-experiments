package main

import (
	"fmt"
	"time"
)

type (
	Action string
	Stage  string
)

type Event struct {
	ActionLabel Action
}

type ApplicationsPerCluster map[string][]string

var deploy Action = "do-deploy"

func (a Action) String() string {
	return string(a)
}

func NewEvent(a Action) *Event {
	return &Event{
		ActionLabel: a,
	}
}

func main() {
	list := make(ApplicationsPerCluster)
	list["one"] = []string{"leo", "luz"}

	fmt.Printf("list: %v", list)

}

func print(message string) {
	fmt.Printf("message: %s at %s", message, time.Now().Format(time.RFC822))
}
