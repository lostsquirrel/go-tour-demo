package main

import (
	"fmt"
	"testing"
)

func TestConstantsiota(t *testing.T) {
	type Grades int

	const (
		A Grades = iota
		B
		C
		D
		E
		F
	)
	fmt.Print(A, B, C, D, E, F)
}
