package main

import (
	"fmt"
)

const (
	a = 'A'
	b = iota
	c
	d = iota
)

func main() {
	fmt.Println(a, b, c, d)
}
