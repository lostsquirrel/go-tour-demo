package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"fmt"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle  {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x, y int) color.Color {
	a := uint8(x ^ y)
	fmt.Println(a)
	return color.RGBA{a, a, a, a}
}
func main() {
	m := Image{}
	pic.ShowImage(m)
}
