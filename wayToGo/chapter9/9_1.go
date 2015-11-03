package main

import (
	"fmt"
	"container/list"
)
func main() {
	l := list.New()
	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)
	
	for e := l.Front(); e!=nil; e=e.Next(){
		fmt.Println(e.Value)
	}
}

