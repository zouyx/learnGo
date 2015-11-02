package main

import (
	"fmt"
)
func main() {
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one":1, "two":2}
	mapCreated:=make(map[string]float32)
	mapAssigned=mapLit

	mapCreated["key1"]=4.5
	mapCreated["key2"]=3.14159
	mapAssigned["two"]=3

	fmt.Printf("Map lit at \"one\" is :%d \n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is :%f \n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is :%d \n", mapLit["two"])
	fmt.Printf("Map assigned at \"ten\" is :%d \n", mapLit["ten"])

	mapNew := new(map[string]float32)
	mapMake := make(map[string]float32)
	fmt.Println("mapNew:",mapNew)
	fmt.Println("mapMake:",mapMake)
}

