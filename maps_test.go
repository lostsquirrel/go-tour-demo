package main

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func TestMap(t *testing.T) {
	var m map[string]common.FloatVertex
	m = make(map[string]common.FloatVertex)
	m["Bell Labs"] = common.FloatVertex{40.68433, -74.39967}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
