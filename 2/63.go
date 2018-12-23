package main

import (
	"fmt"
)

type InSlice struct {
	ptr      *int
	len, cap int
}

func main() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) //append the slice x
	fmt.Println(x)      //"[1 2 3 4 5 6 1 2 3 4 5 6]"
}
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	//...expand z to at least zlen...
	copy(z[len(x):], y)
	fmt.Println(zlen)
	return z
}
