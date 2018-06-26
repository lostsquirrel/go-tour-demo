package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func walk(t *tree.Tree, c chan int)  {
	if t != nil {
		walk(t.Left, c)
		c <- t.Value
		walk(t.Right, c)
	}
}
func Walk(t *tree.Tree, c chan int)  {

	walk(t, c)
	close(c)
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	//fmt.Println(len(c1), len(c2))

	for  {
		e, ok := <-c1
		e2, ok2 := <-c2
		fmt.Println(e, e2, ok, ok2)
		if e != e2 || ok != ok2{
			return false
		}
		if !ok || !ok2 {
			break
		}
	}
	return true
	
}

func main() {
	c := make(chan int)
	t := tree.New(3)
	go Walk(t, c)

	for x := range c {
		fmt.Println(x)
	}
	//fmt.Println(Same(tree.New(1), tree.New(1)))
	//fmt.Println(Same(tree.New(1), tree.New(2)))
	//fmt.Println(Same(tree.New(3), tree.New(2)))
}
