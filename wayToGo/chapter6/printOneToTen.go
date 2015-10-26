package main

import (
	"fmt"
)
func main() {
	print(10)
}

func print(n int) {
	if n < 1{
		return 
	}
	fmt.Println(n)
	n--
	print(n)
	
	return
}
