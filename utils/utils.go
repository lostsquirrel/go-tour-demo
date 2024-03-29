package utils

import (
	"fmt"
	"testing"
)

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func PrintSliceWithDescription(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)

}

func IsEqualsInt(expected, actual int, t *testing.T) {
	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
