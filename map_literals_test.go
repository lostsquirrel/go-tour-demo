package main

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func TestMapLiterals(t *testing.T) {
	var m = map[string]common.FloatVertex{
		"Bell Labs": common.FloatVertex{
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	var n = map[string]common.FloatVertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
	fmt.Println(n)
}
