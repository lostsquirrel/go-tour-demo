package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(time.Second)
	c := make(chan int)
	select {
	case <-c:
		fmt.Println("c")
	case <-timer.C:
		fmt.Println("timeout")
	}
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Println(<-ticker.C)
	}

}
