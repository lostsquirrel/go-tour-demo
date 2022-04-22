package method

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func Scale(v *common.FloatVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestMethodPointer(t *testing.T) {
	v := common.FloatVertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
