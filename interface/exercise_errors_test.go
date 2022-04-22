package iface_demo

import (
	"fmt"
	"math"
	"testing"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		var err ErrNegativeSqrt = ErrNegativeSqrt(x)
		return 0, err
	} else {
		return math.Sqrt(x), nil
	}

}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func TestError2(t *testing.T) {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
