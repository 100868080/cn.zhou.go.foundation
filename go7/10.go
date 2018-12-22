package main

import (
	"fmt"
)

func main() {
	//f1(a,b,c int)int
	var x, y int = 6, 9
	sum := f1(x, y)
	fmt.Println(sum)
}
func f1(a, b int) {
	return a + b
}
