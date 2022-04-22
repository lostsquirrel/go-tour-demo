package main

import (
	"fmt"
	"testing"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func TestNamedResults(t *testing.T) {
	fmt.Println(split(28))
}
