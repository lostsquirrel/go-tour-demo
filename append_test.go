package main

import (
	"go-tour-demo/utils"
	"testing"
)

func TestAppend(t *testing.T) {
	var s []int
	utils.PrintSlice(s)

	s = append(s, 0)
	utils.PrintSlice(s)

	s = append(s, 1)
	utils.PrintSlice(s)

	s = append(s, 2, 3, 4)
	utils.PrintSlice(s) // len=5 cap=6 [0 1 2 3 4] ?
}
