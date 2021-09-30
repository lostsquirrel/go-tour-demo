package go_basic

import (
	"log"
	"testing"
)

func TestAssignArrayToArray(t *testing.T) {
	var a1 [5]string
	a2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	a1 = a2
	log.Printf("%p,%v", &a1, a1)
	log.Printf("%p,%v", &a2, a2)
	a1[0] = "Red222"
	log.Printf("%p,%v", &a1, a1)
	log.Printf("%p,%v", &a2, a2)
}

func TestAssignArrayInDifferentType(t *testing.T) {
	var a1 [4]string
	a2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	log.Printf("%p,%v", &a1, a1)
	log.Printf("%p,%v", &a2, a2)
	//a1 = a2 //Cannot use 'a2' (type [5]string) as the type [4]string
}

func TestAssignArrayPointer(t *testing.T) {
	var a1 [3]*string
	a2 := [3]*string{new(string), new(string), new(string)}

	*a2[0] = "Red"
	*a2[1] = "Blue"
	*a2[2] = "Green"

	a1 = a2
	printString3PArray(a1)
	printString3PArray(a2)
	*a1[0] = "Red 111111111111"
	printString3PArray(a1)
	printString3PArray(a2)
}

func printString3PArray(a [3]*string)  {
	for _, x := range a {
		log.Printf("%s", *x)
	}
	log.Println()
}