package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSwitch(t *testing.T) {
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
