package iface_demo

import "testing"

func TestInterfaceSatisfiedImplicitly(t *testing.T) {

	var i I = &T{"hello"}
	i.M()
}
