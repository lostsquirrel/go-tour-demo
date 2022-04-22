package main

import (
	"fmt"
	"testing"
)

func TestShortVariableDeclaration(t *testing.T) {
	var i, j = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
