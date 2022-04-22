package main

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	var pow []int
	pow = append(pow, 1, 2, 4, 8, 16, 32, 64, 128)
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}
