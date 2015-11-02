package main

import (
	"fmt"
)
func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["new"]=55
	map1["bj"]=20
	map1["washington"]=25

	value, isPresent=map1["bj"]
	if isPresent {
		fmt.Printf("the value of bj is :%d \n", value)
	} else{
		fmt.Printf("map1 doesn't contain bj")
	}

	value, isPresent=map1["paris"]
	fmt.Printf("is Paris in map1? : %t\n", isPresent)
	fmt.Printf("the value is :%d \n", value)

	delete(map1,"washington")
	value, isPresent=map1["washington"]
	if isPresent {
		fmt.Printf("the value of washington is :%d \n", value)
	} else{
		fmt.Printf("map1 doesn't contain washington")
	}
}

