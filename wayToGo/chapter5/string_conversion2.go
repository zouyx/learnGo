package main

import (
	"fmt"
	"strconv"
	"os"
)

func main() {
	var orig string = "abc"
	var an int  
	var newS string  

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, err := strconv.Atoi(orig)
	if err != nil{
		fmt.Printf("orig %s is not an integer - exiting with error \n", orig)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("The int is : %d\n", an)
	an = an + 5 
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is : %s\n", newS)
}
