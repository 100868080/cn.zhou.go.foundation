package main

import (
	"fmt"
)

type rea struct {
	width  int
	height int
}

func (recv rea) area() int {
	return recv.width * recv.height
}
func main() {
	r := rea{456, 312}
	g := rea{456, 31}
	fmt.Println(r.area(), g.area())
}
