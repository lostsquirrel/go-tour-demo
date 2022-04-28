package main

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func TestStructFields(t *testing.T) {
	v := common.IntVertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
