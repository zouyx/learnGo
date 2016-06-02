package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("hello")

	var j = 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	for {
		fmt.Println(j)
		j++
		if j > 100 {
			break
		}
	}
}
