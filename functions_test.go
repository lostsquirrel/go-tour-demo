package main

import (
	"fmt"
	"testing"
)

func add(x, y int) int {
	return x + y
}

func TestFunctions(t *testing.T) {
	fmt.Println(add(26, 37))
}
