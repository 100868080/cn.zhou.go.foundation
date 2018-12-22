package main

import (
	"fmt"
)

func main() {
	sum, div, l, k := f1(6, 5)
	fmt.Println(sum, div, l, k)
}
func f1(a, b int) (sum, div, k int, l float32) {
	//return a + b, a * b, float32(a) / float32(b), a ^ b
	sum = a + b
	div = a * b
	l = float32(a) / float32(b)
	k = a - b
	return
}
