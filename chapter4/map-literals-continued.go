package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell": {
		40.68, -74.39,
	},
	"Google": {
		37.42, -122.08,
	},
}

func main() {
	fmt.Println(m)
	
}
