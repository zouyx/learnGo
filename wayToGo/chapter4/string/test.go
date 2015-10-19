package main

import (
	"fmt"
)

func main() {
	i := uint(0)
	fmt.Println(i)
	//j := ^i
	var j uint = 2147483647
	fmt.Printf("%v, %T", j, j)
	a := j >> 63
	fmt.Println(a)
	b := 32 << a
	fmt.Println(b)
}
