package main

import (
	"fmt"
)
func main() {
	p2 := add2()
	fmt.Printf("call add2 for 3 gives : %v \n", p2(3))

	twoAdder := adder(2)
	fmt.Printf("the result is : %v \n", twoAdder(3))
}

func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(b int) int{
	return func(b int) int{
		return a + b
	}
}
