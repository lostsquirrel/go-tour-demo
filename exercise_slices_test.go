package main

func Pic(dx, dy int) [][]uint8 {

	var res [][]uint8
	for x := 0; x <= dx; x++ {
		each := make([]uint8, dx+1)
		for y := 0; y <= dy; y++ {
			//each[y] = uint8((x + y) / 2)
			//each[y] = uint8(x * y)
			each[y] = uint8(x ^ y)
		}
		res = append(res, each)
	}
	return res
}
