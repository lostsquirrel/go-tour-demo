package method

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func ScaleFunc(v *common.FloatVertexP, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func TestIndirection(t *testing.T) {
	v := common.FloatVertexP{3, 4}
	v.Scale(10)
	ScaleFunc(&v, 10)

	p := &common.FloatVertexP{5, 12}
	p.Scale(10)
	ScaleFunc(p, 10)
	fmt.Println(v, p)

}
