package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	RMB
	EUR
)

var curr = []string{USD: "美元", RMB: "人民币", EUR: "欧元"}

func getCurrency(c Currency) string {
	return curr[c]
}

func main() {
	fmt.Println(getCurrency(USD))
}
