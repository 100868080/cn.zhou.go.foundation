package main

import (
	"fmt"
)

type san struct {
	c int
	k int
	g int
}

func main() {

	r := san{6, 2, 8}
	s := san{654, 546, 8989}
	fmt.Println(a(r), a(s))
}
func a(l san) int {
	return l.c * l.k * l.g / 2
}
