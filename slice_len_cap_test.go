package main

import (
	"go-tour-demo/utils"
	"testing"
)

func TestSliceLenCap(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}
	utils.PrintSlice(s)

	s = s[:0]
	utils.PrintSlice(s)

	s = s[:4]
	utils.PrintSlice(s)

	s = s[2:]
	utils.PrintSlice(s)

	//s = s[6:] // panic: runtime error: slice bounds out of range
	//printSlice(s)
}
