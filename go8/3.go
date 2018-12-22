package main

import (
	"fmt"
)

func main() {
	type student struct {
		id   int
		name string
		card bool
	}
	var s1 = student{133, "liming", false}
	var s2 = &student{123, "pk", true}
	var s3 = student{card: true, id: 564}
	var s4 = &student{name: "king", id: 9845}
	fmt.Println(s1, "\t", s2, "\t", s3, "\t", s4)
}
