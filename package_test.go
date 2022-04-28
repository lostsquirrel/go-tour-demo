package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPackageImport(t *testing.T) {
	//rand.Seed(0)
	fmt.Println("My favorite number is ", rand.Intn(10))
}
