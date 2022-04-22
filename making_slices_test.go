package main

import "testing"

func TestMakingSlice(t *testing.T) {
	a := make([]int, 5)
	printSliceWithDescription("a", a)

	b := make([]int, 0, 5)
	printSliceWithDescription("b", b)

	c := b[:2]
	printSliceWithDescription("c", c)

	d := c[2:5]
	printSliceWithDescription("d", d)
}
