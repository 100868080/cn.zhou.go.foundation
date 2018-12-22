package main

import (
	"fmt"
)

type student struct {
	name string
	id   int
	do   int
	co   bool
}
type date struct {
	year  int
	month int
}

func main() {
	s1 := new(student)
	s1.id = 12345
	s1.name = "li ming"
	fmt.Println(s1)
	var s2 student
	s2.name = "wanggang"
	s2.id = 789456
	fmt.Println(s2)
}
