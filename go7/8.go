package main

import (
	"fmt"
)

func f1(a int) {
	fmt.Scanf("%d", &a)
}

func f2(a, b int) int {
	return a + b
}
func f3(z int) {
	fmt.Println(z)
}
func main() {
	var a int
	var b int = 6
	var x, y int
	f1(a)
	f2(x, y)
	var z int = f2(a, b)
	f3(z)
}
