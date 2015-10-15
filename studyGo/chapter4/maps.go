package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var mm map[string]int

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])


	mm = make(map[string]int)
	mm["joe"] = 123
	fmt.Println(mm["joe"])
	
}
