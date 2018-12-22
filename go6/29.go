package main

import (
	"fmt"
)

func main() {
	var a = []int{'a', 'b'}
	fmt.Println(a)
	var n = []byte{'A', 'B'}
	fmt.Println(n)
	var s = n[:]
	fmt.Println(s)
	var c uint = 'h'
	fmt.Println(c)
}
