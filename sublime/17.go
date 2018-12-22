package main

import (
	. "fmt"
	"math"
)

type ret struct {
	width, height float64
}

type cir struct {
	radius float64
}

func (r ret) area() float64 {
	return r.width * r.height
}

func (c cir) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	c1 := ret{654, 6}
	c2 := ret{2, 3}
	r := cir{645}
	Println(c1.area(), c2.area(), r.area())

	type mon map[string]int
	m := mon{
		"one": 1,
		"two": 2,
	}
	Println(m)
}
