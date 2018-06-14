package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	var res [][]uint8
	for x := 0; x <= dx; x++ {
		each := make([]uint8, dx)
		for y := 0; y <= dy; y++ {
			//each[y] = uint8((x + y) / 2)
			//each[y] = uint8(x * y)
			each[y] = uint8(x ^ y)
		}
		res = append(res, each)
	}
	return res
}

func main() {
	pic.Show(Pic)
}
