package go_basic

import (
	"fmt"
	"log"
	"testing"
)

func print2Dx2(arr [][2]int) {
	for r, row := range arr {
		for c, col := range row {
			log.Printf("%d,%d -> %d", r, c, col)
		}
	}
}
func Test2DInit(t *testing.T) {
	var arr [4][2]int
	print2Dx2(arr[:])
}

func Test2DInitVal(t *testing.T) {
	arr := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	print2Dx2(arr[:])
}

func Test2DInitPart(t *testing.T) {
	arr := [4][2]int{1: {20, 21}, 3: {40, 41}}
	print2Dx2(arr[:])
}

func Test2DInitPart2(t *testing.T) {
	arr := [4][2]int{1: {0: 20}, 3: {1: 41}}
	print2Dx2(arr[:])
}

func Test2DAccess(t *testing.T) {
	var arr [2][2]int
	arr[0][0] = 10
	arr[0][1] = 20
	arr[1][0] = 30
	arr[1][1] = 40
	print2Dx2(arr[:])
}

func Test2DAssign(t *testing.T) {
	var a1 [2][2]int
	var a2 [2][2]int
	a2[0][0] = 10
	a2[0][1] = 20
	a2[1][0] = 30
	a2[1][1] = 40

	a1 = a2
	print2Dx2(a1[:])
	print2Dx2(a2[:])
}

func Test2DAssign2(t *testing.T) {

	var a2 [2][2]int
	a2[0][0] = 10
	a2[0][1] = 20
	a2[1][0] = 30
	a2[1][1] = 40

	var a3 [2]int = a2[1]
	printIntSlices(a3[:])
	var v int = a2[1][0]
	fmt.Println(v)
}

func foo(a [1e6]int) {
	fmt.Sprintln(len(a))
}

func bar(a *[1e6]int)  {
	fmt.Sprintln(len(a))
}
func BenchmarkArrayPass(b *testing.B) {
	var a [1e6]int
	for i := 0; i < b.N; i++ {
		foo(a)
	}
}

func BenchmarkArrayPointerPass(b *testing.B) {
	var a [1e6]int
	for i := 0; i < b.N; i++ {
		bar(&a)
	}
}