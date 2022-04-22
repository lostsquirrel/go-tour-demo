package iface_demo

import (
	"testing"
)

func TestInterfaceValueWithNil(t *testing.T) {

	var i I
	var j *T
	i = j
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}
