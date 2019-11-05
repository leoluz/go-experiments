package main

import "fmt"

type Speaker interface {
	Say(message string)
}

type FmtSpeaker struct{}

func (s *FmtSpeaker) Say(message string) {
	fmt.Println(message)
}

func SayHi(s Speaker) {
	s.Say("hi")
}

func main() {
	s := new(FmtSpeaker)
	SayHi(s)
}
