package main

import "fmt"

func main() {
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
