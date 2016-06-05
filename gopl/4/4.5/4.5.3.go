package main

import (
	"fmt"
	"time"
)

func main() {
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(b)

	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
}
