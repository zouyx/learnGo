package main

import (
	"fmt"
)

func main() {
	var p = getNum()

	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(&p)

}

func getNum() *int {
	v := 1
	fmt.Println(v)
	fmt.Println(&v)
	return &v
}
