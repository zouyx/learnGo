package main

import (
	"fmt"
)

func main() {
	var str string 
	var strP *string 
	if str == "" {
		fmt.Println(`str is ""`)
	}

	if len(str) == 0{
		fmt.Println("str len is 0 ")
	}
	if *strP == "" {
		fmt.Println(`str is ""`)
	}

	if len(*strP) == 0{
		fmt.Println("str len is 0 ")
	}
}
