package main

import (
	. "fmt"
)

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	var tom person
	tom.name, tom.age = "Tom", 23

	bob := person{"paul", 546}

	jim := person{"jimi", 63}

	a, b := Older(tom, bob)
	c, d := Older(tom, jim)
	e, f := Older(bob, jim)
	Println(a, b, "\n", c, d, "\n", e, f)

}
