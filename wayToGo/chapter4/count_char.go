package main

import (
	"fmt"
	"unicode/utf8"
)

func CountChar(str string) (int, int) {
	b := []byte(str)
//	int bytes = len(b)
	var bytes int = len(b)
	var runes int = utf8.RuneCount(b)
	return bytes, runes
}

func main() {
	var i1, j1 int  = CountChar("asSASA ddd dsjkdsjs dk")
	fmt.Println(i1, j1)
	var i2, j2 int = CountChar("asSASA ddd dsjkdsjsこん dk")
	fmt.Println(i2, j2)
	
	
}
