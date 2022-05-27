package main

import (
	"fmt"
	"testing"
	"time"
)

func TestStopGoroutine(t *testing.T) {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done channel is triggered, exit")
				return
			default:
				fmt.Println("working")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(time.Second * 9)
	close(done)
	time.Sleep(time.Second * 3)
}
