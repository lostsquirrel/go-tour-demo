package method

import (
	"fmt"
	"go-tour-demo/common"
	"math"
	"testing"
)

func Abs(v common.FloatVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func TestMethdFunc(t *testing.T) {
	v := common.FloatVertex{3, 4}
	fmt.Println(Abs(v))
}
