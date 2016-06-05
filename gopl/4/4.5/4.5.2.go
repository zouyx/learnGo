package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan string)
	str := "33333"

	fmt.Printf("begin str s:%s ,mem:%p\n", str, &str)
	for i := 0; i < 3; i++ {
		fmt.Printf("loop :%d \n", i)
		go changeStr(str, i, ch)
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("output chan s:%s ,mem:%p \n", <-ch)
	}

	fmt.Printf("final str s:%s ,mem:%p\n", str, &str)
}

func changeStr(s string, i int, ch chan<- string) {
	fmt.Printf("chan s:%s %s ,mem:%p\n", s, &s)
	s = strconv.Itoa(i)
	fmt.Printf("fin chan s:%s ,mem:%p\n", s, &s)

	ch <- s
	return
}
