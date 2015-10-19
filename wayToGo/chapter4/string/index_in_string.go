package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I am Marc, Hi."
	fmt.Printf("The position of Marc is :")
	fmt.Printf("%d \n ", strings.Index(str, "Marc"))

	fmt.Printf("The position of the first instance of hi is :")
	fmt.Printf("%d \n ", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of hi is: ")
	fmt.Printf("%d \n ", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of Burger is :")
	fmt.Printf("%d \n ", strings.Index(str, "Burger"))
}
