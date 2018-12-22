package main

import (
	"fmt"
)

type re struct {
	width  int
	height int
}

func area(s re) int {
	return s.width * s.height
}
func main() {
	r1 := re{4, 3}
	r2 := re{30, 15}
	fmt.Println(area(r1), area(r2))
}
