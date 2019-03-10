package main

import "fmt"

type Celsius float64
type IDnum int

func main() {
	var temp Celsius = 14.1
	var pid IDnum = 9527
	fmt.Print(temp, pid)
}
