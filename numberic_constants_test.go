package main

import (
	"fmt"
	"testing"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func TestNumbericConstants(t *testing.T) {
	//fmt.Println(needInt(Big), )
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Small))
}
