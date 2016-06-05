package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string = "112312321"
	var y string = "1112312321"
	// var z string = "2"

	fmt.Println(x[1:5])
	fmt.Println(x[1:])
	fmt.Println(x[:])
	if x == y {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if x == "1" {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	s := "left foot"
	t := s
	s += ", right foot"

	fmt.Printf("t:%s \n", t)
	fmt.Printf("s:%s \n", s)

	ch := make(chan string)
	str := "33333"
	for i := 0; i < 2; i++ {
		fmt.Printf("loop :%d \n", i)
		go changeStr(&str, &i, ch)
	}

	for i := 0; i < 2; i++ {
		fmt.Printf("output chan s:%s \n", <-ch)
	}

	fmt.Printf("final str s:%s \n", str)
}

func changeStr(s *string, i *int, ch chan<- string) {
	fmt.Printf("chan s:%s \n", *s)
	str := strconv.Itoa(*i)
	*s = str
	fmt.Printf("fin chan s:%s \n", *s)

	ch <- *s
	return
}
