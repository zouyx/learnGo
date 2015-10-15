package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("version: %s", runtime.Version())
}
