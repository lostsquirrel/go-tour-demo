package common

import "math"

type FloatVertex struct {
	X, Y float64
}

func (v FloatVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
