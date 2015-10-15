package main

import (
	"fmt"
	"math/rand"
)

// init method is called before main
func init(){
	fmt.Println("init !")
}

// main method is a core method for this program
func main() {
	fmt.Println("hello wolrd!",rand.Intn(10))
}
