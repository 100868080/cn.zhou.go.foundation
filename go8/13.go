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
		*people
		der string
	}
	t1 := teacher{&people{"zhengzhi", true}, "computer science"}
	fmt.Println(t1, t1.der, t1.name, "\t", t1.people, "\t", t1.sex)
}
