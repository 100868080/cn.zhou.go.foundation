package main

import (
	"fmt"
)

type abc struct {
	id   int
	name string
}

func (abc) my() {
	fmt.Println("gogogogogogogogo!")
}
func main() {
	a := abc{}
	a.my()
	//fmt.Println(a.my())
}
