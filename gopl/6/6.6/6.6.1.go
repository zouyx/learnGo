package main

import (
	"fmt"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	f1 := f()
	fmt.Println(f, &f1, f1)
	f1 = f()
	fmt.Println(f, &f1, f1)
	f1 = f()
	fmt.Println(f, &f1, f1)
	f1 = f()
	fmt.Println(f, &f1, f1)

	f = nil
	f1 = 0

	f = squares()
	f1 = f()
	fmt.Println(f, &f1, f1)

}
