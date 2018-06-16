package main

import (
	"math"
	"fmt"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		f = - f
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}