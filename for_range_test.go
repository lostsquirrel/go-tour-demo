package main

import (
	"fmt"
	"testing"
)

func TestForRangeString(t *testing.T) {
	for index, char := range "TestForRangeString" {
		fmt.Println(index, char)
	}
}

func TestForRangeMap(t *testing.T) {
	for k, v := range map[string]int{"a": 1, "b": 2, "c": 3} {
		fmt.Println(k, v)
	}
}

func TestForRangeArrya(t *testing.T) {
	for i, v := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i, v)
	}
}
