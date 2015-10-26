package main

import (
	"fmt"
)
func main() {
	//var s []string = {"a","b","c","d"}
	s := []string {"a","b","c","d"}
	printEachElement(s...)
}

func printEachElement(s ...string) {
	if len(s) == 0 {
		return
	}
	for i, v := range s {
		fmt.Printf("index:%d , value: %s \n",i,v)
	}
}
