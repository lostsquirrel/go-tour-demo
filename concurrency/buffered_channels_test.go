package main

import (
	"fmt"
	"testing"
)

func TestBufferedChannels(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
