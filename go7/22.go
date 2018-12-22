package main

import (
	"fmt"
)

func main() {
	f1(1, 2, 3, 4, 5, 6, 7)
}
func f1(a ...int) {
	f2(a...)
	f2(a[2:]...)
}
func f2(a ...int) {
	fmt.Println(a)
}
