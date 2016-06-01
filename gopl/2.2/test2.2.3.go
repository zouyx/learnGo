package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("begin.")
	//練習 1.1： 脩改echo程序，使其能夠打印os.Args[0]。
	fmt.Println(os.Args[0])
	//練習 1.2： 脩改echo程序，使其打印value和index，每個value和index顯示一行。
	for index, value := range os.Args[0:] {
		fmt.Print(index)
		fmt.Println("," + value)
	}

	//練習 1.3： 上手實踐前面提到的strings.Join和直接Println，併觀察輸出結果的區别。
	fmt.Println(strings.Join(os.Args[0:], " "))
}
