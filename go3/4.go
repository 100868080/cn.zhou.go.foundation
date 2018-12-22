package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	var x float32
	fmt.Println("3 ge shu:")
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println("-x_")
	x = float32(a+b+c) / 3.0
	fmt.Println("%6.2f", x)
}
