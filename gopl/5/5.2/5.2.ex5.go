package main

import (
	"fmt"
)

//練習 4.5： 寫一個函數在原地完成消除[]string中相鄰重複的字符串的操作。
func main() {
	str := "jjalskkdjffllas"
	bs := []byte(str)
	// var bsNew []byte = make([]byte, len(str), len(str))
	var bsNew []byte = []byte{}
	for i := 0; i < len(str)-1; i++ {
		if i == 0 {
			bsNew = append(bsNew, bs[i])
		}

		if bs[i] != bs[i+1] {
			bsNew = append(bsNew, bs[i+1])

		}
	}
	strNew := string(bsNew)
	fmt.Println(strNew)
	fmt.Println(bsNew)
}
