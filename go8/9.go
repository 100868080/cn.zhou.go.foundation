package main

import (
	"fmt"
)

func main() {
	type ab struct {
		name string
		id   int
		nan  bool
	}
	type cd struct {
		ab
		id int
	}
	t := cd{ab{"loin", 4564, true}, 4654987465}
	fmt.Println(t)
}
