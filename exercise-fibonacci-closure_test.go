package main
import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib, fibB1, fibB2 := 0, 0, 0
	return func() int {
		if fibB2 == 0 {
			fibB2 = 1
		} else {
			fib = fibB1 + fibB2
			fibB2 = fibB1
			fibB1 = fib
		}
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}