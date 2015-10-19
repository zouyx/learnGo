package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi there !"
	var newS string

	newS = strings.Repeat(str, 3)

	fmt.Printf("new string is %s ", newS)
}
