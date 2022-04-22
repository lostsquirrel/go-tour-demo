package main

import (
	"fmt"
	"math"
	"testing"
)

func Sqrt(x float64) float64 {
	var res float64
	var epsilon = 1e-9
	for z := 1.0; math.Abs(x-z*z) > epsilon; {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z, x)
		res = z
	}
	return res
}

func TestLoopsAndFunctions(t *testing.T) {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(4))
}
