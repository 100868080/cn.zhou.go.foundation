package main

import (
	"fmt"
)

type coor struct {
	x int
	y int
}

func main() {
	r := coor{56, 456}
	r.wc()
	fmt.Println(r)
}
func (vg *coor) wc() {
	vg.x, vg.y = vg.y, vg.x
	fmt.Println(vg)
}
