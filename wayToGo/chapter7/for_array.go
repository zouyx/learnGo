package main

import (
	"fmt"
)
func main() {
	var arr [5]int
	
	for i:=0; i < len(arr); i++ {
		arr[i] = i * 2
	}
	
	for i, v := range arr{ 
		fmt.Printf("index %d is %d\n", i, v)
	}

	a := [...]string{"a", "b", "c", "d"}
for i := range a {
    fmt.Println("Array item", i, "is", a[i])
}
}

