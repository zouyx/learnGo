package main

import (
	"fmt"
)

func main() {
	var x string

	var b bool

	if b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if !b {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

	// var y []string
	if x == "1" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// if y[0] == "1" {
	// 	fmt.Println("true")
	// }
	// fmt.Println("false")
}
