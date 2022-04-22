package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}
