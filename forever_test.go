package main

import (
	"fmt"
	"testing"
)

func TestForever(t *testing.T) {
	for {
		fmt.Println("forever")
	}
}
