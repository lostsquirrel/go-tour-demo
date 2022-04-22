package iface_demo

import (
	"fmt"
	"image"
	"testing"
)

func TestImageRGBA(t *testing.T) {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.ColorModel())
	fmt.Println(m.At(0, 0).RGBA())
}
