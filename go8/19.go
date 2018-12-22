//多个method的定义
package main

import (
	"fmt"
	"math"
)

type re struct {
	width  int
	height int
}
type circle struct {
	radius float32
}

func (rt re) area() int {
	return rt.width * rt.height
}
func (rt circle) area() float32 {
	return rt.radius * rt.radius * math.Pi
}
func main() {
	r := re{231, 5}
	c := circle{6.33}
	fmt.Println(r.area(), c.area())
}
