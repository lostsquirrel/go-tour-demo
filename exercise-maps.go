package main

import (
	"strings"
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
func main() {
	wc.Test(WordCount)
}