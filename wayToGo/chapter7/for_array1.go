package main

import (
	"fmt"
)
const (
	limit int =16
)

func main() {
	var arr [limit]int
	
	for i:=0; i < len(arr); i++ {
		arr[i] = i 
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}
	
}

