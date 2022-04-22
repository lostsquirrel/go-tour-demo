package main

func TestSliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)

	//s = s[6:] // panic: runtime error: slice bounds out of range
	//printSlice(s)
}
