package main

import (
	"fmt"
)

func main() {
	var i int
	var f float64
	var b bool
	var s string
	const c ="%v %v %v %q \n"
	fmt.Printf(c,i,f,b,s)
}
