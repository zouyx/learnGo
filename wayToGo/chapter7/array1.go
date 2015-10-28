package main

import (
	"fmt"
)
func main() {
	var arr [5]int
	arr2 := arr
	
	for i:=0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	for i:=0; i < len(arr2); i++ {
		arr2[i] = i * 3
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	
}

