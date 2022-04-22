package common

import "math"

type FloatVertexP struct {
	X, Y float64
}

func (v *FloatVertexP) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *FloatVertexP) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
