package main

import (
	"fmt"
)

func main() {
	i,j:=42,2701

	p:=&i//point to i 
	fmt.Println("p,point to i:",*p) //red i through the pointer
	
	*p=21//set i throgh the pointer
	fmt.Println("i:",i) //set the new value of i

	p =&j
	*p=*p /37 //divide j through the pointer
	fmt.Println("j:",j) //see the new value of j
}
