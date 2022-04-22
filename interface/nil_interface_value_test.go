package iface_demo

import "testing"

func TestNilInterfaceValue(t *testing.T) {
	var i I
	describe(i)
	// i.M()
}
