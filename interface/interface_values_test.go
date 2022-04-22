package iface_demo

import (
	"math"
	"testing"
)

func TestInterfaceValue(t *testing.T) {
	var i I
	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}
