package main

import (
	"fmt"
	"testing"
)

func TestDeclare(t *testing.T)  {
	var a [3]int
	fmt.Println(a)
}

func TestAssignment(t *testing.T)  {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	if a[0] != 1 && a[2] != 3 {
		t.Fail()
	}
}

func TestInitialValue(t *testing.T)  {
	a := [3]int{1, 2, 3}
	if a[0] != 1 && a[2] != 3 {
		t.Fail()
	}
}

func TestMultiLineInitialValue(t *testing.T)  {
	greetings := [4]string{
		"Good morning",
		"Good afternoon",
		"Good evening",
		"Good night",
	}

	fmt.Println(greetings[3])

}

func TestAutomaticArrayLenghtDeclaration(t *testing.T)  {
	greetings := [...]string{
		"Good morning",
		"Good afternoon",
		"Good evening",
		"Good night",
	}

	fmt.Println(greetings[2])
}

func TestArrayComparion(t *testing.T)  {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 3, 3}
	c := [3]int{3, 3, 3}
	d := [3]int{1, 2, 3}
	if a != d && a == b && b == c && c == d {
		t.Fail()
	}
}

func TestArrayIterationByFor(t *testing.T)  {

	a := [...]int{1, 2, 3, 4, 5}
	for index := 0; index < len(a); index++ {
		fmt.Printf("a[%d] = %d", index, a[index])
	}
	fmt.Println()
}

func TestArrayIterationByRange(t *testing.T)  {
	a := [...]int{1, 2, 3, 4, 5}
	for index, value := range(a) {
		fmt.Printf("a[%d] = %d", index, value)
	}
	fmt.Println()
}

func TestArrayIterationByRangeWithBlank(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}
	for _, value := range(a) {
		fmt.Print(value)
	}
	fmt.Println()
}

func TestMultiDimensionalArrayInitial(t *testing.T) {
	a := [3][2]int {
		[2]int{1, 2},
		[2]int{3, 4},
	}

	fmt.Println(a)
}

func TestMultiDimensionalArrayShotSyntx(t *testing.T) {

	a := [3][2]int {
		{1, 2},
		{3, 4},
		{5, 6},
	}

	fmt.Println(a)
}