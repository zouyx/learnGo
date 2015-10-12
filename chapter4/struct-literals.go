package main

import (
	"fmt"
)

type Vertex struct{
	X,Y int
}

var(
	v1=Vertex{1,2}// type is Vertex
	v2=Vertex{X:1}// Y:0 is miss 
	v3=Vertex{}// Y:0 and X:0
	v4=Vertex{Y:9,X:1}// Y:9 and X:1 
	p=&Vertex{1,2}// type is *Vertex 
)

func main() {
	fmt.Println("v1:",v1,"p:",p,"v2:",v2,"v3:",v3)
	fmt.Println("v4:",v4)
}
