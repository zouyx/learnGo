package main

import (
	"fmt"
)
func main() {
	var arr1 [6]int
	fmt.Println(len(arr1))
	var slice1 []int = arr1[2:len(arr1)]
	
	for i:=0; i < len(arr1); i++ {
		arr1[i] = i 
	}

	fmt.Println(arr1)

	for i:=0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d \n", i ,slice1[i])
	}
	
	fmt.Printf("the length of arr1 is %d \n", len(arr1))
	fmt.Printf("the length of slice1 is %d \n", len(slice1))
	fmt.Printf("the cap of slice1 is %d \n", cap(slice1))

//	slice1 = slice1[:len(slice1)-1]
	slice1 = slice1[2:4]
	for i := 0; i<len(slice1); i++{
		fmt.Printf("slice at %d is %d \n", i ,slice1[i])
	}
	fmt.Printf("the length of slice1 is %d \n", len(slice1))
	fmt.Printf("the cap of slice1 is %d \n", cap(slice1))
}

