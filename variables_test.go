package main

import "fmt"

var c, python, java bool

func TestVariableScope() {
	var i int
	fmt.Println(i, c, python, java)
}
