package main

import (
	"fmt"
)

type people struct {
	name  string
	phone string
}
type teacher struct {
	people
	dear string
}
type student struct {
	people
	school string
}

func (r people) say() {
	fmt.Println("I,m %s you can call me on %s\n", r.name, r.phone)
}
func (s student) say() {
	fmt.Println("I,m %s call me on %s\n", s.name, s.phone)
}
func main() {
	t1 := teacher{people{"zhengzhi", "465-9846556"}, "computer science"}
	t2 := student{people{"li ming", "564+8978898"}, "yela universty"}
	t1.say()
	t2.say()
}
