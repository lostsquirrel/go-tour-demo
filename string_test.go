package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "abcdefghijklmnopqrstuv"
	for c := range s {
		fmt.Printf("%d", c)
		fmt.Println(c)
		fmt.Println(string(rune(c)))
	}
}
