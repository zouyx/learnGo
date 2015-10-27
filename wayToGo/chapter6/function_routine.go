package main

import (
	"fmt"
)
func main() {
	var g int 
	go func(i int){
		s:=0
		for j:=0; j<i; j++{ s+=j }
		g = s
	}(1000)
	fmt.Println(g)
}

