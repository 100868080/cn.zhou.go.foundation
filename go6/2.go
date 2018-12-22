package main

import (
	"fmt"
)

func main() {
	var x = []int{1, 2, 3, 4, 5, 6}
	var i int = 2
	x[i] = x[i+2] * x[i-1]
	fmt.Println(x[i])
}
