package main

import (
	"fmt"
	"testing"
)

func TestStructLiterals(t *testing.T) {
	v1 := IntVertex{1, 2}
	v2 := IntVertex{X: 1}
	v3 := IntVertex{}
	p := &IntVertex{1, 2}
	fmt.Println(v1, v2, v3, p)
}
