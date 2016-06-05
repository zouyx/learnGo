package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)
	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages["alice"])

	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
		"joe":     34,
	}

	fmt.Println(ages)
	ch := make(chan string)

	for i := 0; i < 2; i++ {
		go delKey(&ages, i, ch)
	}

	for i := 0; i < 2; i++ {
		<-ch
	}

	fmt.Println(ages)
}

func delKey(m *map[string]int, num int, ch chan<- string) {
	if num == 1 {
		fmt.Println("del joe")

		delete(*m, "joe")

	} else {
		fmt.Println("del alice")
		delete(*m, "alice")
	}
	ch <- ""
	return
}
