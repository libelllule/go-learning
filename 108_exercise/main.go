package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	y := float64(0)
	for y != z {
		y = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("Math: ", math.Sqrt(2))
}
