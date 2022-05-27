package main

import (
	"fmt"
	"testing"
	"time"
)

func productor(c chan<- int) {
	for {
		c <- 1
		time.Sleep(time.Second)
	}
}

func consumer(c <-chan int) {
	for {
		fmt.Println(<-c)
	}
}

func TestOneWayChannel(t *testing.T) {
	var c = make(chan int)
	go productor(c)
	go consumer(c)
	time.Sleep(time.Second * 5)
}
