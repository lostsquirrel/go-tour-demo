package main

import (
	"fmt"
	"testing"
)

var i, j int = 1, 2

func TestVariable_with_initializer(t *testing.T) {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
