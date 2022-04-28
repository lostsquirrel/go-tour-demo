package main

import (
	"go-tour-demo/utils"
	"testing"
)

func TestMakingSlice(t *testing.T) {
	a := make([]int, 5)
	utils.PrintSliceWithDescription("a", a)

	b := make([]int, 0, 5)
	utils.PrintSliceWithDescription("b", b)

	c := b[:2]
	utils.PrintSliceWithDescription("c", c)

	d := c[2:5]
	utils.PrintSliceWithDescription("d", d)
}
