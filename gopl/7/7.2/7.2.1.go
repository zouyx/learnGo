package main

import (
	"fmt"
)

type Point struct{ X, Y float64 }

// same thing, but as a method of the Point type
func (p *Point) GetDistance(x, y float64) {
	p.X = x
	p.Y = y
}
func main() {
	p := &Point{1, 2}

	p.GetDistance(2, 1)
	(*p).GetDistance(2, 1)
	fmt.Println(*p)
	fmt.Println(*p)

	q := &Point{3, 4}
	q.GetDistance(4, 3)

	fmt.Println(*p)
	fmt.Println(*q)

}
