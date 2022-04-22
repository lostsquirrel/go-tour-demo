package method

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func TestMehodPointerReceiver(t *testing.T) {
	v := &common.FloatVertexP{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
