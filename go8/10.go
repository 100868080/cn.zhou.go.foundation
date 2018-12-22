package main

import (
	"fmt"
)

func main() {
	type people struct {
		name string
		sex  bool
	}
	type teacher struct {
		people
		der string
	}
	type der struct {
		teacher
		id int
	}
	bead := der{}
	bead.id = 907
	bead.name = "liming"
	bead.sex = true
	bead.der = "lioma"
	//bead.people = "name":dad"
	fmt.Println(bead)
}
