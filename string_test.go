package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "abcdefghijklmnopqrstuv"
	for c := range s {
		fmt.Println(c)
	}
}
