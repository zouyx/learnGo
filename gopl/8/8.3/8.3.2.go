package main

import (
	"fmt"
)

type Reader interface {
	Read(args string)
}
type Writer interface {
	Write(args string)
}

type RW struct {
}

func (rw RW) Read(args string) {
	fmt.Println("read string:", args)
}

func (rw RW) Write(args string) {
	fmt.Println("Write string:", args)
}

func main() {
	var w Writer = RW{}
	w.Write("joe")

}
