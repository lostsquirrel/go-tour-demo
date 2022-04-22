package main

import (
	"fmt"
	"testing"
)

func TestMapLiterals(t *testing.T) {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}

	var n = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
	fmt.Println(n)
}
