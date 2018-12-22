//匿名字段方法
package main

import (
	"fmt"
)

type people struct {
	name string
}
type teacher struct {
	people
	depar string
}

func (p people) sayhi() {
	fmt.Printf("I,m %s \n", p.name)
}

type Speaker interface {
	sayhi()
}

func main() {
	peo := people{"zhang san"}
	tea := teacher{people{"zheng zhi"}, "computer science"}
	var is Speaker
	is = &peo
	is.sayhi()
	is = &tea
	is.sayhi()
}
