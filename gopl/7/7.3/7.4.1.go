package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// traditional function
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (c *ColoredPoint) ScaleBy(f float64) {
	c.X *= 2
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	distanceFromP := p.Distance

	fmt.Println(distanceFromP(q.Point)) // "5"

	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
}
