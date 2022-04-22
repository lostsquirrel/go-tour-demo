package main

import (
	"fmt"
	"testing"
)

func TestForIsWhile(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
