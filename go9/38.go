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
	depar string
}
type student struct {
	people
	school string
}

func (r people) sayHi() {
	fmt.Printf("Hi,I'm %s you can call me on %s \n", r.name, r.phone)
}
func main() {
	t1 := teacher{people{"zhengzhi", "312-456897"}, "computer science"}
	s1 := student{people{"liming", "645-987456"}, "yale university"}
	t1.sayHi()
	s1.sayHi()
}
