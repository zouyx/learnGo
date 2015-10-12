package main

import (
	"fmt"
	"math"
)
func Sqrt(x float64) float64{
	z:=1.0
	k:=0.0
	for i:=0;i<10;i++{
		k=z-(math.Pow(z,2)-x/(2*z))
		fmt.Println(k)
		if k<0{
			return z
		}else{
			z=k
		}	
	}
	return z
}

func main() {
	fmt.Println(
		Sqrt(3),
	)
}
