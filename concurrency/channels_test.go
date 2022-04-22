package main

import (
	"fmt"
	"testing"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func TestChannel(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	i := len(s) / 2
	go sum(s[:i], c)
	go sum(s[i:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
