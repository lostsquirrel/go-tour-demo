package method

import (
	"fmt"
	"go-tour-demo/common"
	"math"
	"testing"
)

func AbsFunc(v common.FloatVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestIndirectionValues(t *testing.T) {

	v := common.FloatVertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &common.FloatVertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
