package main

import (
	"fmt"
	"math/rand"
)

func TestPackageImport() {
	//rand.Seed(0)
	fmt.Println("My favorite number is ", rand.Intn(10))
}
