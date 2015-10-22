package main

import (
	"fmt"
)

func main() {
LABEL1:
	for i := 1 ; i <= 5 ; i++ {
		for j :=0 ; j<=5; j++{
			if j ==4 {
				continue LABEL1
			}
			fmt.Println("i is: %d, and j is: %d\n", i, j)
		}
	}
}
