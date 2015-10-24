package main

import (
	"fmt"
	"math"
)

func main() {
	result , err := MySqrt(-1)
	fmt.Printf("result : %f \n", result)
	fmt.Println("error : ", err)
}

func MySqrt(f float64) (returnF float64, err error){
	if f < 0 {
		err = fmt.Errorf("input : %f is invalid!", f)	
		returnF = 0.0
		return
	}
	returnF = math.Sqrt(f)
	return
}

