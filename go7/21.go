package main

import (
	"fmt"
)

func main() {
	var i int
	f1(5, 5, 6, 6, 2, 1, 5)
	f1(6)
	var b = [7]int{1, 5, 6}
	for i = 0; i <= 6; i++ {
		k := b[i]

		fmt.Println(k)
		k += b[i]
		fmt.Println(k)
	}

}
func f1(a ...int) {
	fmt.Println(a)

}
