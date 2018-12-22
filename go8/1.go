package main

import (
	"fmt"
)

func main() {
	type student struct {
		id   int
		call string
	}
	type ad float32
	var v ad = 56.1897778
	fmt.Println(v)
	var s student
	s.id = 120366548
	s.call = "wang"
	fmt.Println(s)
}
