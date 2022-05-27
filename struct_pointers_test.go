package main

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func TestStructPointers(t *testing.T) {
	v := common.IntVertex{X: 1, Y: 2}
	p := &v
	p.X = 1e9
	fmt.Println(v.X)
}
