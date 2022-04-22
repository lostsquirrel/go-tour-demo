package main

import (
	"fmt"
	"testing"
)

func TestStructPointers(t *testing.T) {
	v := IntVertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v.X)
}
