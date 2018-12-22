package main

import (
	"fmt"
)

func main() {
	var m map[string]int
	m = make(map[string]int)
	m["k"] = 1

	fmt.Println(m)
	var b = map[string]int{}
	b["l"] = 3
	fmt.Println(b)
	x, y := m["k"]
	if y {
		fmt.Println(x, y)
	}
	k, ok := m["z"]
	if ok {
		fmt.Println(k, ok)
	} else {
		fmt.Println(k)
	}

	v, ok := b["l"]
	if ok {
		fmt.Println(v, ok)
	}
}
