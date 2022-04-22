package main

import (
	"strings"
	"testing"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, e := range strings.Split(s, " ") {
		v := m[e]
		m[e] = v + 1
	}
	return m
}
func TestMakeMap(t *testing.T) {
	wc.Test(WordCount)
}
