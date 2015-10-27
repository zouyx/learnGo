package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
//	where := func() {
  //  _, file, line, _ := runtime.Caller(1)
    //log.Printf("%s:%d", file, line)
//	}
	var where = log.Print


	var y int
	var x = f()
	where()
	for i := 0; i < 10; i++ {
		y = x(i)
		fmt.Println(y)
	}
}

func f() func(int) int {
	var x int
	return func(i int) int {
		var j = f()
		if i <= 1 {
			x = 1
		} else {
			x = j(i-1) + j(i-2)
		}
		return x
	}
}
