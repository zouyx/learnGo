package main

import (
	"flag"
	"fmt"
)

type MyFlag struct {
}

func (m MyFlag) Set(string) error {
	fmt.Println("string")
	return nil
}

func (m MyFlag) String() string {
	return "string"
}

func main() {
	var flag flag.Value = new(MyFlag)
	flag.Set("joe")
}
