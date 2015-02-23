package main

import "fmt"

type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

type Player struct {
	*User
	GameId int
}

func NewPlayer(id int, name, location string, gameId int) *Player {
	return &Player{
		&User{id, name, location},
		gameId,
	}
}

func main() {
	p := NewPlayer(43, "LeoLuz", "Montreal", 1313)
	fmt.Println(p.Greetings())
	fmt.Printf("My gameId is: %v\n", p.GameId)
}
