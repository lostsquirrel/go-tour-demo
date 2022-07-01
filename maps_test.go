package main

import (
	"fmt"
	"go-tour-demo/common"
	"testing"
)

func TestMapDefine(t *testing.T) {
	var m0 map[string]common.FloatVertex
	fmt.Println(m0)

}
func TestMapDefineAndAccess(t *testing.T) {

	m := make(map[string]common.FloatVertex)
	m["Bell Labs"] = common.FloatVertex{X: 40.68433, Y: -74.39967}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
}

func TestMapFunctioneValue(t *testing.T) {
	myFuncMap := map[string]func() int{
		"x": func() int { return 1 },
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["x"]
	fmt.Println(f())
}

func TestMapAccess(t *testing.T) {
	m := map[string]int{"a": 1}
	v, ok := m["a"]
	if ok {
		fmt.Println(v)
	}
}

func TestMapIterate(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
