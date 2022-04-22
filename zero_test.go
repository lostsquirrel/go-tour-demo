package main

import (
	"fmt"
	"testing"
)

func TestZeroValues(t *testing.T) {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
