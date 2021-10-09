package go_basic

import "testing"

func TestSliceCreate(t *testing.T) {
	s := make([]string, 5)
	printStringSlicesInfo(s)
}

func TestSliceCreateCap(t *testing.T) {
	s := make([]string, 3, 5)
	printStringSlicesInfo(s)
}

func TestSliceCreateCapLessLength(t *testing.T) {
	//s := make([]string, 5, 3) //  len larger than cap in make([]string)
	//printSlicesInfo(s)
}

func TestSliceLiteral(t *testing.T) {
	a := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	printStringSlices(a)
}

func TestSliceIndexPosition(t *testing.T) {
	s := []string{99: "99"}
	printStringSlices(s)
}

func TestEmptySlice(t *testing.T) {
	s1 := make([]string, 0)
	printStringSlicesInfo(s1)
	s2 := []string{}
	printStringSlicesInfo(s2)
	var s3 []string
	printStringSlicesInfo(s3)
}

func TestSliceArrayLiteral(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	printIntSlices(s)
	s[1] = 22
	printIntSlices(s)
}

func TestSliceOfSlice(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	ns := s[1:3]
	printIntSlicesInfo(ns)
	ns[1] = 33
	printIntSlices(ns)
}