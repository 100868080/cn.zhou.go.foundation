package main

import (
	"fmt"
)

func main() {
	var z, x int = 3, 4
	f := sum
	f1(z, x, f)
}
func f1(z, x int, sum func(int, int) int) {
	fmt.Println(sum(z, x))
}
func sum(z, x int) int {
	return z + x
}
