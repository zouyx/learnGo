package main

import (
	"fmt"
	"gopl.io/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	fmt.Println("begin")
	var s = []string{"32", "212"}
	for _, arg := range s[0:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
