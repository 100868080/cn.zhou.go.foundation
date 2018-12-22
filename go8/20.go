package main

import (
	"fmt"
)

type coor struct {
	x int
	y int
}

func (recv coor) swap() {
	var temp int
	temp, recv.x, recv.y = temp, recv.x, recv.x
	fmt.Println(recv)
}
func (recv *coor) sap() {
	var temp int
	temp, recv.x, recv.y = temp, recv.x, recv.x
	fmt.Println(recv)
}
func main() {
	r := coor{12, 32}
	r.swap()
	r.sap()
	fmt.Println(r)
}
