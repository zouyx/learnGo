package main

import (
	"fmt"
	"math/rand"
)

const Pi = 3.1415926
const b string = "abc"
const b1 = "abc"
const beef, two, c = "meat", 2, "veg"
const Monday, Tuesday, Wednesday = 1, 2, 3
const (
	Thrusday, Friday = 4, 5
	Saturday, Sunday =6, 7
)
const (
	Unknown = iota
	Female 
	Male  
)

func init() {
	fmt.Println("init !")
	fmt.Println(Pi,b,b1)
	fmt.Println(beef,two,c)
	fmt.Println(beef,two,c)
	fmt.Println(Monday, Tuesday, Wednesday)
	fmt.Println(Unknown, Female, Male)
}

func main() {
	fmt.Println("hello wolrd!", rand.Intn(10))
}
