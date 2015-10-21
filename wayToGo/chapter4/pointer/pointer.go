package main

import (
	"fmt"
)

func main() {
	var i = 5
	var iP = &i // var iP *int iP = &i
	fmt.Printf("An integer :%d , it's location in mem: %p value :%d \n", i, &i, *iP)
	i = 9
	fmt.Printf("The value has changed.An integer :%d , it's location in mem: %p value :%d \n", i, &i, *iP)
}
