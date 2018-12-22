package main

import (
	"fmt"
)

func main() {
	sum := f1(3, 6)
	fmt.Println(sum)
}
func f1(a, b int) int {
	return a + (a * b)
}
