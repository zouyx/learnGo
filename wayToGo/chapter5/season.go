package main

import (
	"fmt"
)

func main() {
	k := 5
	switch {
		case k == 1, k == 2, k ==12 :
			fmt.Println("Winter")
		case k >=3 && k < 6:
			fmt.Println("Spring")
		case k >=6 && k < 8:
			fmt.Println("Summer")
		case k >=9 && k < 11:
			fmt.Println("Autumn")
		default:
			fmt.Println("Num is invalid")
	}
}
