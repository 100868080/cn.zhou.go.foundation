package main

import (
	"fmt"
)

func f1() {
	fmt.Println("the godern is sky!")
}

func f2(a, b int, c string) {
	//return a + b
	fmt.Println(a, b, c)
}

func f3(a, b int) int {
	return a + b
}

func main() {
	var a, b int = 6, 9
	var c string = "golor"
	f1()
	f2(a, b, c)

	sum := f3(a, b)
	fmt.Println(sum)
}
