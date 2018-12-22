package main

import (
	"fmt"
)

type sd struct {
	l int
	k int
}

func (cd sd) re() int {
	return cd.l * cd.k
}
func main() {
	t := sd{4, 645}
	fmt.Println(t.re())
}
