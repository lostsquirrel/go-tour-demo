package main

import (
	"fmt"
	"testing"
	"time"
)

func say(s string) {

	for i := 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(s)
	}
}

func TestGoroutine(t *testing.T) {
	go say("world")
	say("hello")
}
