package main

//嵌套map定义
import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m = map[string]map[string]Vertex{
}

var mm = map[string]Vertex{
	"Labs": {40.11,-32.11},
}

func main() {
	m["Bell"]= mm
	
	fmt.Println(m)
	
}
