package main

import (
	"fmt"
	"math/cmplx"
	"testing"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func TestBasicTypes(t *testing.T) {
	fmt.Printf("Type %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value %v\n", z, z)
}
