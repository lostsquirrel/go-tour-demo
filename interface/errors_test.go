package iface_demo

import (
	"fmt"
	"testing"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %v", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

func TestError(t *testing.T) {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
