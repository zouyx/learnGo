package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int{
	words := strings.Fields(s) 
	fmt.Println(words)
	
	wordCount := map[string]int{} 
	for _,word := range words{
		if  wordCount[word] == 0 {
			wordCount[word] = 1
		}else {
			wordCount[word] += wordCount[word] 
		}
		
	} 

	return wordCount
}

func main() {
//	fmt.Println(mm["joe"])
	fmt.Println(WordCount("i am joe joe"))
	
}
