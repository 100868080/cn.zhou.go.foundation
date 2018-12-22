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

type emp struct {
	Human
	com string
}

func (h *Human) say() {
	fmt.Printf("%s call me %s\n", h.name, h.phone)
}

func (e *emp) say() {
	fmt.Printf("I %s ,work at %s,call:%s\n", e.name, e.com, e.phone)
}
func main() {
	mark := Student{Human{"mark", 25, "132-89564-adzx"}, "MIT"}
	sam := emp{Human{"sam", 45, "8998-523222-yyuu"}, "golang inc"}
	mark.say()
	sam.say()
}
