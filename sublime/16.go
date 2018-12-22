package main

import (
	. "fmt"
)

type ret struct {
	width, height float64
}

func area(r ret) float64 {
	return r.height * r.width
}

func main() {
	x := ret{21, 21}
	y := ret{6, 6}
	Println(area(x), area(y))
}
