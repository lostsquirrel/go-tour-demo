package main

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	var appleNum int
	fmt.Printf("Number of apples?")
	num, err := fmt.Scan(&appleNum)
	fmt.Print(num, err, appleNum)
}
