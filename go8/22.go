package main

import (
	"fmt"
)

type object struct {
	id   int
	name string
}

func (object) ms() {
	fmt.Println("this is aobject!")
}
func (*object) msg() {
	fmt.Println("this is a!")
}
func main() {
	ob := object{}
	p := &ob
	ob.ms()
	p.msg()
}
