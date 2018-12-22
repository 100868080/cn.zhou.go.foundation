package main

import (
	"fmt"
)

func main() {
	var q = [5]int{1, 2, 3, 4}
	var r = q[3:]
	var w = q[0:]
	var n int
	n = copy(w, r)
	fmt.Println(n, w)
	fmt.Println(q)
	fmt.Println(w)
}
