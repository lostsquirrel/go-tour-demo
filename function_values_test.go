package main

import (
	"fmt"
	"math"
	"testing"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func TestFunctionValues(t *testing.T) {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
