package main

import (
	"fmt"
)
func main() {
	fv := func (){
		fmt.Println("hello world!")
	}
	
	fv()

	str := "hello world1"
	func (s string){
		fmt.Println(s)
	}(str)

}


