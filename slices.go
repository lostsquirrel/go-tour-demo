package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}

	var a []int = primes[1:4]
	fmt.Println(a)
}
