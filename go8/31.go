package main

import (
	"fmt"
)

type student struct {
	num   int
	name  string
	sex   bool
	age   int
	score int
}

func main() {
	var s1 student

	var s = []int{1, 2, 3, 4, 6}
	fmt.Println(s)
	s1.name = "zio"
	fmt.Println(s1)
	s1 = student{123456, "zhou", false, 55, 98}
	s2 := student{64545665, " asda", true, 45, 64}
	fmt.Println(s1, "\n", s2)
	//fmt.Println(s2)
}
