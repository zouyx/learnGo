package main

import (
	"fmt"
)
func main() {
	newI :=new([]int)
	fmt.Printf("type :%T \n",newI)
	makeI :=make([]int,0)
	fmt.Printf("type :%T \n", makeI)
}

