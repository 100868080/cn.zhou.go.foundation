package main

import (
	"fmt"
)

func main() {
	var a = [3][5]int{{1, 2, 3}, {4, 5}}
	fmt.Println(a)

	var p [5]int

	var z = p[1:]
	var k = []int{2, 2, 323, 15, 5, 554}
	var h = make([]int, 8, 26)m
	fmt.Println(p)

	fmt.Println(z)
	fmt.Println(k)

	fmt.Println(h)
	for q, w := range p {
		fmt.Println(q, w)
	}
	var t = append(z, 56, 659, 98)
	fmt.Println(t)


}
