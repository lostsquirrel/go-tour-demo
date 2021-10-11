package go_basic

import (
	"fmt"
	"testing"
)

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
	printIntSlices(s)
}

func TestSliceOutOfRange(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	ns := s[1:3]
	//ns[3] = 44  // index out of range
	printIntSlices(ns)
}

func TestSliceAppend(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	ns := s[1:3]
	ns = append(ns, 60)
	printIntSlices(s)
	printIntSlices(ns)
}

func TestSliceAppendLengthAndCapacity(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	printIntSlicesInfo(s)
	s = append(s, 60)
	printIntSlicesInfo(s)
	
}

func TestSliceThreeIndex(t *testing.T) {
	s := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	sx := s[2:3:4]
	printStringSlicesInfo(sx)
	//sxx := s[2:3:6]
	//printStringSlicesInfo(sxx)
//	 slice bounds out of range
	s3 := s[2:3:3]
	printStringSlicesInfo(s3)
	s3 = append(s3, "Kiwi")
	printStringSlicesInfo(s3)
	printStringSlices(s)
	printStringSlices(s3)
}

func TestSliceAppendAnotherSlice(t *testing.T) {
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))

}

func TestSliceIterating(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	for index, value := range s {
		fmt.Printf("Index: %d Value: %d \n", index, value)
	}
}

func TestSliceIteratingCopy(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	for index, value := range s {
		fmt.Printf("Value: %d Value-Addr: %X Elem-Addr: %X \n", value, &value, &s[index])
	}
}

func TestSliceIteratingIgnoreIndex(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	for _, value := range s {
		fmt.Printf("Value: %d \n",  value)
	}
}

func TestSliceIteratingFor(t *testing.T) {
	s := []int{10, 20, 30, 40, 50}
	for i :=2; i < len(s); i++ {
		fmt.Printf("Index: %d Value: %d \n", i, s[i])
	}
}
