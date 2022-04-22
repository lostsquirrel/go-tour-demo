package main

import (
	"fmt"
	"testing"
)

func TestRangeC(t *testing.T) {

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
