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
	e := NewEvent(deploy)
	print(e.ActionLabel.String())
}

func print(message string) {
	fmt.Printf("message: %s at %s", message, time.Now().Format(time.RFC822))
}
