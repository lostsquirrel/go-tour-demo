package method

import (
	"fmt"
	"math"
	"testing"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		f = -f
	}
	return float64(f)
}

func TestMethodContined(t *testing.T) {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
