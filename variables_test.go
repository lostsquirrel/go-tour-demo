package main

import (
	"fmt"
	"testing"
)

var c, python, java bool

func TestVariableScope(t *testing.T) {
	var i int
	fmt.Println(i, c, python, java)
}
