package main

import (
	"fmt"
	"math"
)

type Abser interface{
	Abs() float64
}

type Vertex struct{
	X,Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v. X + v.Y * v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f //a MyFloat 实现了Abser
	a = &v //a *vertex 实现了 Abser 

	// 下面一行，v是一个vertex（而不是 ＊Vertex）
	// 所以没有实现Abser
	a = v 

	fmt.Println(a.Abs())
}
