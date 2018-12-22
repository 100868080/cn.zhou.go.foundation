package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
}
type ele struct {
	Human
	company string
}

func (h *Human) sayhi() {
	fmt.Printf("hi,%s I,m %s\n", h.name, h.phone)
}
func main() {
	mark := Student{Human{"mark", 25, "231-546"}, "mit"}
	sam := ele{Human{"sam", 45, "456-645"}, "golang inc"}

	mark.sayhi()
	sam.sayhi()
}
