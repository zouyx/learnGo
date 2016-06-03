package main

import (
	"fmt"
)

func main() {
	p := new(int)                // p, *int 類型, 指向匿名的 int 變量
	fmt.Printf("type :%T \n", p) // type
	fmt.Println(*p)              // "0"
	*p = 2                       // 設置 int 匿名變量的值爲 2
	fmt.Println(*p)              // "2"

	pj := new(int)
	q := new(int)

	fmt.Println(pj)
	fmt.Println(q)
	fmt.Println(pj == q) // "false"

	rt := delta(1,
		2,
	)
	fmt.Println(rt)
}

func delta(old, new int) int { return new - old }
