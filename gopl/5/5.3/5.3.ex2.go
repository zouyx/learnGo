package main

import (
	"fmt"
)

func main() {
	str := "jjalskkdjffllas"
	wordCount := map[string]int{}
	b := []byte(str)

	for i := 0; i < len(b)-1; i++ {
		s := string(b[i])
		wordCount[s]++
	}

	fmt.Println(wordCount)
}
