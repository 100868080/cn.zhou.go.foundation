package main

import (
	"fmt"
)

type recatangle struct {
	width  int
	height int
}

func (recv recatangle) area() int {
	return recv.width * recv.height
}
func main() {
	r1 := recatangle{645, 123}
	r2 := recatangle{456, 123}
	fmt.Println(r1.area(), r2.area())
}
