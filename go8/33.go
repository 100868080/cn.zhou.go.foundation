package main

import (
	"fmt"
)

type hk struct {
	width  int
	height int
}

func (tc hk) link() int {
	return tc.width * tc.height
}
func main() {
	tc := hk{56, 456}
	fmt.Println(tc.link())
}
