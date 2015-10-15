package main

import (
	"fmt"
)

func d(){
	defer fmt.Println("joe")
}

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	d()
}
