package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func TestMutex(t *testing.T) {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 5; i++ {
		go c.Inc("somekey")

	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x)
		}(i)
	}
	wg.Wait()
}
