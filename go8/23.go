package main

import (
	"fmt"
)

type people struct {
	id    string
	phone string
}
type teacher struct {
	people
	dear string
}

func (r people) sayhi() {
	fmt.Println("I'm %s name you can call me %c:\n", r.id, r.phone)
}
func main() {
	t1 := teacher{people{"zhengzhi", "010-22002"}, "computer science"}
	t2 := teacher{people{"liming", "231-8456645"}, "yale university"}
	t1.sayhi()
	t2.sayhi()
}
