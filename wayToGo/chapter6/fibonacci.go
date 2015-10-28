package main

import (
	"fmt"
	"time"
)
func main() {
	res := 0
	start := time.Now()	
	for i := 0; i<=40 ; i++{
		res = fibonacci(i)
		fmt.Printf("f(%d) is %d \n ", i, res)
	}
	end := time.Now()	
	fmt.Println("delta:",end.Sub(start))
}

func fibonacci(n int) (res int) {
	if n <= 1{
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}
