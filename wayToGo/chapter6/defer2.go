package main

import (
	"fmt"
)
func main() {
	f()
}

func f(){
	for i:=0; i<5; i++{
		defer fmt.Printf("%d ",i)
	}
}
