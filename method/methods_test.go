package method

import (
	"fmt"
	"testing"

	"go-tour-demo/common"
)

func TestMethd(t *testing.T) {
	v := common.FloatVertex{3, 4}
	fmt.Println(v.Abs())
}
