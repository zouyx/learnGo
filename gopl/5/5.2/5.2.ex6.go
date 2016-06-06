package main

import (
	"fmt"
	"unicode"
)

// 練習 4.6： 編寫一個函數，原地將一個UTF-8編碼的[]byte類型的slice中相鄰的空格（參考unicode.IsSpace）替換成一個空格返迴
func main() {
	str := "j  jkdafa  kda j"
	bs := []byte(str)
	// var bsNew []byte = make([]byte, len(str), len(str))
	var bsNew []byte = []byte{}
	for i := 0; i < len(str)-1; i++ {
		if i == 0 {
			bsNew = append(bsNew, bs[i])
		}
		var r1 rune = rune(bs[i])
		var r2 rune = rune(bs[i+1])
		if !(unicode.IsSpace(r1) && unicode.IsSpace(r2)) {
			bsNew = append(bsNew, bs[i+1])

		}
	}
	strNew := string(bsNew)
	fmt.Println(strNew)
	fmt.Println(bsNew)
}
