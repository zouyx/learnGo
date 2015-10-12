package main

import (
	"fmt"
)

type ErrNegativeSqrt struct{
	Num float64
}

func (e *ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number:%v ", float64(e.Num))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, &ErrNegativeSqrt{x}
	} 
	return x, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
