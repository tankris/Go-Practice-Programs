package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; z != math.Sqrt(x) && i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%g -> %g\n", z, math.Sqrt(x))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
