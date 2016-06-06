package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID, Salary, ManagerID   int
	Name, Address, Position string
	DoB                     time.Time
	age                     int
}

var joe Employee = Employee{Name: "joe"}

func main() {
	fmt.Println(joe)
	joe.Salary = 100
	fmt.Println(joe)

	position := &joe.Position
	*position = "coach"
	fmt.Println(joe)

	will := &joe
	fmt.Println(*will)

	// em := getStru()

	// fmt.Printf("%p\n", &em)

	var e Employee = Employee{}
	setStru(e)
	fmt.Printf("%p %v\n", &e, e)

	var e Employee = Employee{}

}

// func getStru() Employee {
// 	var e Employee = Employee{}
// 	fmt.Printf("%p\n", &e)
// 	return e
// }

func setStru(em Employee) {
	em.ID = 1
	fmt.Printf("%p %v\n", &em, em)
}
