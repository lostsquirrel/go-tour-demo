package main

import (
	"fmt"
	"testing"
)

type Celsius float64
type IDnum int

func TestTypeDeclaration(t *testing.T) {
	var temp Celsius = 14.1
	var pid IDnum = 9527
	fmt.Print(temp, pid)
}
