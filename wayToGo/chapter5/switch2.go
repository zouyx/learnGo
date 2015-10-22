package main

import (
	"fmt"
)

func main() {
	var num1 int = 7
	
	switch {
		case num1 < 0 :
			fmt.Println("Num is negative")
		case num1 > 0 && num1 < 10 :
			fmt.Println("Num is between 0 and 10")
		default:
			fmt.Println("Num is 10 or greater")
	}
}
