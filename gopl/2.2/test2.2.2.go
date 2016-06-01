package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = "  "
	}

	fmt.Println(s)

	var str string
	var i int8
	var j int16
	var ui uint8
	var uj uint16

	fmt.Println(str)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(ui)
	fmt.Println(uj)

}
