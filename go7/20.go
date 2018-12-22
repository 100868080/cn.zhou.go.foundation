package main

import (
	"fmt"
)

func main() {
	f1(1, 2, 3)
	f1(1, 2, 3, 4, 5, 6, 7)
	//	fmt.Println(arg)
}
func f1(arg ...int) {
	var r int
	for _, arg := range arg {
		fmt.Println(arg)

		r += arg
		fmt.Println(arg)

	}

	fmt.Println(r)

	fmt.Println(arg)
}
