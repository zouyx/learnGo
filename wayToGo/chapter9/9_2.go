package main

import (
	"fmt"
	"unsafe"
)
func main() {
	var l int64 = 100
	fmt.Println(unsafe.Sizeof(l))
}

