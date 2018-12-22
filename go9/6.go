package main

import (
	"fmt"
	"strconv"
)

type human struct {
	name  string
	age   int
	phone string
}

//通过这个方法human实现了fmt.stringer
func (h human) string() string {
	return "<" + h.name + "-" + strconv.Itoa(h.age) + "year-@" + h.phone + "<"
}
func main() {
	bob := human{"bob", 39, "-132564-897456xxx"}
	fmt.Println("this human is:", bob)
}
