package main

import "fmt"

type Point struct {
	X, Y int
}

var (
	p = Point{1, 2}
	q = &Point{1, 2}
	r = Point{Y: 1}
	s = Point{}
)

type Bootcamp struct {
	Latitude  float64
	Longitude float64
}

func main() {
	x := new(Bootcamp)
	y := &Bootcamp{}
	fmt.Println(*x == *y)
	fmt.Println(p.X, q.Y, r, s)
}
