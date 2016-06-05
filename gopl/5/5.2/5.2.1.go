package main

import (
	"fmt"
	"strconv"
)

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	curr := []string{"1", "2", "3", "4", "5"}
	fmt.Println(curr)
	// reverse(curr)

	reverse(curr[:2])
	reverse(curr[2:])
	reverse(curr)
	fmt.Println(curr)

	var i []string
	fmt.Println(i == nil)
	fmt.Println(len(i) == 0)
	reverse(i)
	fmt.Println(i)
	i = []string{}
	fmt.Println(i == nil)
	fmt.Println(len(i) == 0)

	for j := 0; j < 10; j++ {
		// i = append(i, strconv.Itoa(j), strconv.Itoa(j))
		i = append(i, strconv.Itoa(j))
		fmt.Printf("%d cap=%d len=%d %v\n", j, cap(i), len(i), i)
	}

	top := i[len(i)-1]
	fmt.Println(top)

	i = i[:len(i)-1]
	fmt.Println(i)
}
