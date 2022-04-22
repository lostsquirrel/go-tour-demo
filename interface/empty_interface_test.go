package iface_demo

import (
	"go-tour-demo/utils"
	"testing"
)

func TestEmptyInterface(t *testing.T) {
	var i interface{}
	utils.Describe(i)

	i = 42
	utils.Describe(i)

	i = "hello"
	utils.Describe(i)
}
