package main

import (
	"fmt"
)

var x int = 1

func main() {
	fmt.Println(x)
	x := "hello!"

	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}

	fmt.Println("") // "255 0 1"

	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // "255 0 1"

	var i int8 = 127
	fmt.Println(i, i+2, i*i) // "127 -128 1"
}
