package main

import (
	"fmt"
	"go-tour-demo/utils"
	"testing"
)

func TestSlice(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13}

	var a []int = primes[1:4]
	fmt.Println(a)
}

func TestSliceNew(t *testing.T) {
	mySlice := new([]int)
	fmt.Printf("%T\n", mySlice)
}

func TestSliceMake(t *testing.T) {
	mySlice := make([]int, 0)
	fmt.Printf("%T\n", mySlice)
	utils.PrintSlice(mySlice)
}
func TestSliceMake2(t *testing.T) {
	mySlice := make([]int, 10)
	fmt.Printf("%T\n", mySlice)
	utils.PrintSlice(mySlice)
}
func TestSliceMake3(t *testing.T) {
	mySlice := make([]int, 10, 20)
	fmt.Printf("%T\n", mySlice)
	utils.PrintSlice(mySlice)
}

func TestSliceSize(t *testing.T) {
	a := []int{}
	b := []int{1, 2, 3}
	c := a
	fmt.Printf("%p %p %p\n", a, b, c)
	d := append(b, 1)
	fmt.Printf("%p %p %p\n", a, b, d)
}

func TestSliceValueModify(t *testing.T) {
	mySlice := []int{1, 2, 3, 4, 5}
	for _, value := range mySlice {
		value *= 2
	}
	utils.PrintSlice(mySlice)
	for index := range mySlice {
		mySlice[index] *= 2
	}
	utils.PrintSlice(mySlice)
}
