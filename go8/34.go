package main

import (
	"fmt"
)

type seek struct {
	gao  int
	kuan int
}

func (ng seek) jisuan() int {
	return ng.gao * ng.kuan
}
func main() {
	rc := seek{456, 239999991}
	fmt.Println(rc.jisuan())
}
