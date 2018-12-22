//接口的转换
package main

import (
	"fmt"
)

type people struct {
	name string
	age  int
}
type student struct {
	people
	school string
}
type peopleinfo interface {
	getpeopleinfo()
}
type studentinfo interface {
	getpeopleinfo()
	getstudentinfo()
}

func (p people) getpeopleinfo() {
	fmt.Println(p)
}
func (s student) getstudentinfo() {
	fmt.Println(s)
}
func main() {
	var is studentinfo = student{people{"li ming", 18}, "yale universty"}
	is.getstudentinfo()
	is.getpeopleinfo()
	var ip peopleinfo = is
	ip.getpeopleinfo()
}
