package main

import (
	"testing"
)

func TestNewPerson(t *testing.T) {
	person := NewPerson("Bugs", "Bunny")
	if person.Firstname != "Bugs" {
		t.Error("Firstname should be Bugs, got:", person.Firstname)
	}
	if person.Lastname != "Bunny" {
		t.Error("Lastname should be Bunny, got:", person.Lastname)
	}
}
