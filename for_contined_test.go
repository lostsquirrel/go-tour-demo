package main

import (
	"fmt"
	"testing"
)

func TestForC(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
