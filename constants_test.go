package main

import (
	"fmt"
	"testing"
)

const Pi = 3.14

func TestConstants(t *testing.T) {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
