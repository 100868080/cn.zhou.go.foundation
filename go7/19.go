package main

import (
	"fmt"
)

func main() {
	x, y := f1(99, 56)
	fmt.Println(x, y)
}
func f1(a, b int) (x, y int) {
	x = a + b
	y = a * b
	return
}
