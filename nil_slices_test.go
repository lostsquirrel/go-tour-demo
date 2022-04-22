package main

import (
	"fmt"
	"testing"
)

func TestNilSlice(t *testing.T) {
	var a []int
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("nil!")
	}
}
