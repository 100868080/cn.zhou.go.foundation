package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	var i int
	//defer fmt.Println(a)
	for i = 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println(a)
}
