package main

import (
	"fmt"
)

type people struct {
	name string
	sex  bool
}
type teacher struct {
	*people
	dear string
}

func main() {
	t1 := teacher{&people{"zhengzhi", false}, "computer science"}
	fmt.Println(t1, t1.name, t1.dear, t1.sex)
}
