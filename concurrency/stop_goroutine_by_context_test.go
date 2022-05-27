package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func process(ctx context.Context, timeout time.Duration) {
	workTimeout := time.NewTimer(timeout)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context triggered exit")
			return
		case <-workTimeout.C:
			fmt.Println("work timeout")
			return
		default:
			fmt.Println("working")
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func TestStopGoroutineByContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	go process(ctx, 100*time.Millisecond)
	<-ctx.Done()
	fmt.Println("main:", ctx.Err())
}
