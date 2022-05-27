package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestRangAndClose(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func TestRangeBuffer(t *testing.T) {
	c := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Printf("send: %d\n", n)
			c <- n
		}
		close(c)
	}()
	for v := range c {
		fmt.Printf("recive: %d\n", v)
	}

}
