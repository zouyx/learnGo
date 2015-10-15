package main

import (
	"fmt"
)

func fibonacci() func() int{
	a := 0
	b := 1
	sum := 0
	return func() int{
		sum = a + b	
		a = b
		b = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0;i <10 ; i++{
		fmt.Println(f())
	}
}
