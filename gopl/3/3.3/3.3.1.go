package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(&x)

	getNum(p)

	fmt.Println(x)
	fmt.Println(*p)
}

func getNum(p *int) {
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

	*p += 1

	*p++
}
