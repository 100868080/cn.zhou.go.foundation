package main

import (
	"fmt"
)

func f1(a, b int) int {
	return a + b
}
func f2(c string) {
	fmt.Println(c)
}
func main() {
	var c string = "the china town!"
	var a, b int = 8, 6
	f2(c)
	r := f1(a, b)
	fmt.Println(r)
}
