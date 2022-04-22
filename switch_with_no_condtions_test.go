package main

import (
	"fmt"
	"testing"
	"time"
)

func TestSwitchWithNoCondtions(t *testing.T) {
	n := time.Now()
	switch {
	case n.Hour() < 12:
		fmt.Println("Good morning")
	case n.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
