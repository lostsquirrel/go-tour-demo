package main

import "fmt"

func sum(a []int, c chan int)  {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	i := len(s) / 2
	go sum(s[:i], c)
	go sum(s[i:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x + y)
}