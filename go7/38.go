package main

import (
	"fmt"
)

func main() {
	f1(1, 2, 3, 4)
	f1(4, 5, 6, 7, 8, 9)
}
func f1(max ...int) {
	var Max, i int
	Max = max[i]
	for _, v := range max {
		if Max < v {
			Max = v

		}
	}
	fmt.Println(Max)
	return
}
