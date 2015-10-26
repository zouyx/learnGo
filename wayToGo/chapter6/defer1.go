package main

import (
	"fmt"
)
func main() {
	fmt.Println(a())
}

func a() (i int){
	i = 0
	defer fmt.Println(i)
	i++
	return
}

