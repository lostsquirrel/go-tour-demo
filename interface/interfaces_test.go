package iface_demo

import (
	"fmt"
	"math"
	"testing"
)

type Abser interface {
	Abs() float64
}
type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.Y*v.Y + v.X*v.X)
}

func TestInterface(t *testing.T) {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f

	a = &v

	//a = v  //Vertex does not implement Abser (Abs method has pointer receiver)
	fmt.Println(a.Abs())
}
