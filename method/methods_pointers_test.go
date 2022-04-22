package method

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func TestMethodPointer2(t *testing.T) {
	v := common.FloatVertexP{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
