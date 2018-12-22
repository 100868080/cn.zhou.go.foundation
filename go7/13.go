package main

import (
	"fmt"
)

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var s []int = a[:]
	f1(s)
	fmt.Println(s)
}
func f1(s []int) {
	s[0] += 6
	s[3] += 5
	fmt.Println(s)
}
