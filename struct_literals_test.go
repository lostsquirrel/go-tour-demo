package main

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func TestStructLiterals(t *testing.T) {
	v1 := common.IntVertex{1, 2}
	v2 := common.IntVertex{X: 1}
	v3 := common.IntVertex{}
	p := &common.IntVertex{1, 2}
	fmt.Println(v1, v2, v3, p)
}
