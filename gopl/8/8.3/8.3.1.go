package main

import (
	"fmt"
)

type ReadWriter interface {
	Read(args string)
}

type RW struct {
}

func (rw RW) Read(args string) {
	fmt.Println("read string:", args)
}

func (rw *RW) Write(args string) {
	fmt.Println("Write string:", args)
}

func main() {
	var any interface{}
	any = "hello"
	fmt.Println(any)
	any = 1
	fmt.Println(any)
}
