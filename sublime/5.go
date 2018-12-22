package main

import (
	"fmt"
	//"os"
)

func main() {
	a := 3 + 5.1i
	fmt.Println(a)
	fmt.Println(real(a), "+", imag(a), "=", real(a)+imag(a))
	var (
		b        = 6
		c        = "sting"
		d        = false
		g string = "language"
	)
	fmt.Println(b, c, d, g)

	const (
		f = 65
		h = ".string."
	)

	const (
		//j  = 6
		a1 = iota
		a2 = iota
		a3 = iota
	)
	fmt.Println(a1, a2, a3)

	fmt.Println(f, h, g+h, h+g)

}
