package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years) ", p.Name, p.Age)
}

func main() {
	a := Person {"joe", 26}
	z := Person {"michelle", 29}
	fmt.Println(a,z)
}
