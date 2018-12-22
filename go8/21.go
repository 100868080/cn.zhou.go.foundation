package main

import (
	"fmt"
)

type coor struct {
	x int
	y int
}

func (recv *coor) swap() {
	var t int
	t, recv.x, recv.y = t, recv.x, recv.y
	fmt.Println(recv)
}
func main() {
	r := coor{90, 32}
	p := &r
	p.swap()
	fmt.Println(r)
}
