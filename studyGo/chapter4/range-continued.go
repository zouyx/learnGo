package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
		fmt.Printf("index = %d \n", i)
	}

	for _,value := range pow {
		fmt.Printf("value = %d \n", value)
	}
}
