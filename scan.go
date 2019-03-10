package main

import "fmt"

func main() {
	var appleNum int
	fmt.Printf("Number of apples?")
	num, err := fmt.Scan(&appleNum)
	fmt.Print(num, err, appleNum)
}
