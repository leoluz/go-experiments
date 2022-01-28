package main

import "fmt"

type Petable interface {
	Speak() string
	Name() string
	Play()
}
type Pet struct {
	name string
}

type Dog struct {
	Petable
	Breed string
}

func (p *Pet) Play() {
	fmt.Println(p.Speak())
}

func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

func (p *Pet) Name() string {
	return p.name
}

func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Speak(), d.Breed)
}

func main() {
	d := Dog{Petable: &Pet{name: "spot"}, Breed: "pointer"}
	fmt.Println(d.Name())
	// this will crash because Go doesn't support method overload
	//fmt.Println(d.Speak())
	d.Play()
}
