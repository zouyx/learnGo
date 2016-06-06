package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	//匿名函数,叼
	w.X = 1
	w.Y = 2
	w.Circle.Point.Y = 3

	fmt.Printf("%#v", w)
}
