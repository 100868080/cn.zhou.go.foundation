package main

import (
	"fmt"
)

func main() {
	var a = [5]int{1, 2}
	fmt.Println(a)
	b := []int{5, 6, 2}
	fmt.Println(b)
	double := [2][4]int{[4]int{1, 2, 5}, [4]int{32}}
	fmt.Println(double)
	c, d, e := 's', "d", 'e'

	fmt.Println(c, d, e, c+e)
	f := []byte{'s', 'f'}
	fmt.Println(f, f[1], f[0:len(f)])

	g := [3]int{1, 5, 6}
	h := g[:]
	fmt.Println(g, h)
	h = []int{6}
	i := 65
	fmt.Println(h, g, cap(h), cap(g))
	h = append(h, i)
	b := copy(b, h)
	fmt.Println(h, b)
}
