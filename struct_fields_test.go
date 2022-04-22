package main

import (
	"fmt"
	"testing"
)

func TestStructFields(t *testing.T) {
	v := IntVertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
