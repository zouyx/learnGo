package main

import (
	"fmt"
)

func main() {
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("here is the pointer p : %p \n", p)
	fmt.Printf("here is the string *p : %s \n", *p)
	fmt.Printf("here is the string s : %s \n", s)
}
