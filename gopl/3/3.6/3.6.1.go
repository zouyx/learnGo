package main

import (
	"fmt"
	"tempconv"
)

func init() {
	fmt.Println("init....")
}

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! -273.15°C"

	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"

	//練習 2.1： 向tempconv包添加類型、常量和函數用來處理Kelvin絶對溫度的轉換，Kelvin 絶對零度是−273.15°C，Kelvin絶對溫度1K和攝氏度1°C的單位間隔是一樣的。

}
