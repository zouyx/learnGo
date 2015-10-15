package main

import (
	"fmt"
	"math/cmplx"
)
var(
	toBe bool =false
	maxInt uint64 =1<<64-1
	z complex128 =cmplx.Sqrt(-5+12i)
)

func main() {
	const f="%T(%v)\n"
	fmt.Printf(f,toBe,toBe)
	fmt.Printf(f,maxInt,maxInt)
	fmt.Printf(f,z,z)
}
