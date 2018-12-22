package main

import (
	"fmt"
)

func f3(a int) {
	a += 1
	fmt.Println(a)
}
func main() {
	var b = [5]int{1, 2, 3, 4, 5}
	f3(b[2])
	fmt.Println(b)

	f2(b)
}
func f2(g [5]int) {
	g[0] += 2999999999999999999
	fmt.Println(g)
}
