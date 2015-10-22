package main

import (
	"fmt"
)
const (
	FIZZ_BUZZ = 15
	FIZZ = 3
	BUZZ = 5
)

func main() {
	for i := 1 ; i <=100 ;i++{
		switch {
			case i % FIZZ_BUZZ ==0 : 
				fmt.Println("FizzBuzz")
			case i % FIZZ == 0 :
				fmt.Println("Fizz")
			case i % BUZZ == 0 :
				fmt.Println("Buzz")
			default :
				fmt.Println(i)
		}
	}
}
