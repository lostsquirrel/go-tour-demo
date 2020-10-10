package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlag = flag.Int("n", 1234, "help message for flag n")
	var bFlag = flag.Bool("b", false, "bool flag 1")
	flag.Parse()
	//fmt.Printf("Type: %T Value: %v\n", nFlag, nFlag)
	fmt.Println("n = ", *nFlag)
	fmt.Println("b = ", *bFlag)
	//fmt.Printf("n = %d", *nFlag)
}