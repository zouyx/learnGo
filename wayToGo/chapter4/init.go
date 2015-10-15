package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("init !")
}

func main() {
	fmt.Println("hello wolrd!", rand.Intn(10))
}
