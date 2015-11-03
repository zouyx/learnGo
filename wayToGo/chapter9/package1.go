package main

import (
	"fmt"
	"./pack1"
)
func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Println("str:",test1)
	fmt.Println("int:",pack1.Pack1Int)
}

