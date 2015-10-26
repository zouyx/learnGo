package main

import (
	"fmt"
)
func main() {
	res := 0
	for i := 0; i<=10 ; i++{
		res = fibonacci(i)
		fmt.Printf("f(%d) is %d \n ", i, res)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1{
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}
