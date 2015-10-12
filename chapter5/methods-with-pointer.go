package main

import (
	"fmt"
	"math"
)

type Vertex struct{
	X,Y float64
}

func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	fmt.Println(v)
	return math.Sqrt(v.X * v. X + v.Y * v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())

	b := Vertex{3, 4}
	b.Scale(5)
	fmt.Println(b, b.Abs())
}
