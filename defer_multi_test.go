package main

import (
	"fmt"
	"testing"
)

func TestMultiDefer(t *testing.T) {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
