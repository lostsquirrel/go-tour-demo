package main

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {

	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
}
