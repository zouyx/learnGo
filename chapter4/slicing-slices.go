package main

import (
	"fmt"
)

func main() {
	p :=  []int{2,3,5,7,13}
	fmt.Println("p == ",p)
	fmt.Println("p[1:4] == ",p[1:4])

	//省略下表表示从0开始
	fmt.Println("p[:3] == ",p[:3])

	//省略上标表示到len（s）结束
	fmt.Println("p[2:] == ",p[2:])

	for i:= 0; i < len(p); i++{
		fmt.Printf("p[%d] == %d \n",i,p[i])
	}
}
