package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"joe":     34,
		"bob":     21,
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%v %v \n", name, ages[name])
	}

	var agesNil map[string]int
	fmt.Println(agesNil == nil)
	fmt.Println(len(agesNil) == 0)
	// agesNil["a"] = 1

	age, ok := ages["mic"]
	fmt.Printf("%v %v \n", age, ok)

}
