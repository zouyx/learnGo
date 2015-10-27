package main

import (
	"fmt"
)
func main() {
	var f = adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func adder() func(int) int{
	var x int
	return func(delta int) int{
		x += delta
		return x
	}
}
