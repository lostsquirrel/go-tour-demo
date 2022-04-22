package main

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("world")
	fmt.Println("hello")
}
