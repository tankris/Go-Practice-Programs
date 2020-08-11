package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: ")
}

func (e error) Error() string {
	return fmt.Sprintf("")
}

func Sqrt(x float64) (error, float64) {

	if x < 0 {
		return ErrNegativeSqrt(x), x
	}

	z := 1.0
	for i := 0; z != math.Sqrt(x) && i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("%g -> %g\n", z, math.Sqrt(x))
	}

	return nil, z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
