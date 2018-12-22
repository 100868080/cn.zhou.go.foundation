package main

import (
	"fmt"
)

func main() {
	f := c(99)
	fmt.Println(f(3))
	fmt.Println(f(1))
}
func c(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
