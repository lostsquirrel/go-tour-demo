package main

import (
	"runtime"
	"fmt"
)

func main() {
	var osName string
	switch os := runtime.GOOS; os {
	case "darwin":
		osName = "OS X"
	case "linux":
		osName = "Linux"
	default:
		osName = os
	}
	fmt.Printf("Go runs on %s.", osName)
}
