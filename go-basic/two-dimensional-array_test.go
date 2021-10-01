package go_basic

import (
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