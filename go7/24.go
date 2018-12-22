package main

import (
	"fmt"
)

func main() {
	f := func(a, b int) int {
		return a + b
	}
	sum := f(2, 3)
	fmt.Println(sum)
	sum = func(a, b int) int {
		return a + b
	}(87, 56)
	fmt.Println(sum)
}
