package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("this os is : %s \n ", goos)
	path := os.Getenv("PATH")
	fmt.Printf("path is : %s ", path)
}
