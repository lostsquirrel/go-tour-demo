package main

import (
	"flag"
	"fmt"
)

func main() {
	var nFlag = flag.Int("n", 1234, "help message for flag n")
	var bFlag = flag.Bool("b", false, "bool flag")
	var fFlag = flag.Float64("f", 0, "float 64 flag")
	var sFlag = flag.String("s", "foo", "string flag")
	var sVar string
	flag.StringVar(&sVar, "sv","svar", "string var flag")
	flag.Parse()
	//fmt.Printf("Type: %T Value: %v\n", nFlag, nFlag)
	fmt.Println("n = ", *nFlag)
	fmt.Println("b = ", *bFlag)
	fmt.Println("f = ", *fFlag)
	fmt.Println("s = ", *sFlag)
	fmt.Println("sv = ", sVar)
	fmt.Println(len(flag.Args()), "tail:", flag.Args())
}