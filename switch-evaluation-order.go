package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	var res string
	switch time.Saturday {
	case today + 0:
		res = "Today"
	case today + 1:
		res = "Tomorrow"
	case today + 2:
		res = "In two days"
	default:
		res = "Too far away"

	}
	fmt.Println(res)
}
