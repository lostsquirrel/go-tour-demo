package main

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
