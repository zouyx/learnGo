package main

import (
	"fmt"
)

type ReadWriter interface {
	Read(args string)
}

// func (rw ReadWriter) Write(args string) {
// 	fmt.Println("write string:", args)
// }

type RW struct {
}

func (rw RW) Read(args string) {
	fmt.Println("read string:", args)
}

func (rw *RW) Write(args string) {
	fmt.Println("Write string:", args)
}

func main() {
	// rw2 := new(ReadWriter)
	// rw2.Read("joe")

	rw := new(RW)
	rw.Read("joe")
	rw.Write("joe")

	// RW{}.Write("joe2")
	// RW{}.Read("joe2")
	var rw2 RW = RW{}
	rw2.Read("joe2")

	rw1 := &RW{}
	rw1.Read("joe")
}
