package main

import (
	"fmt"
	"math"
	"testing"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func TestSqrt(t *testing.T) {
	fmt.Println(sqrt(2), sqrt(-4))
}
