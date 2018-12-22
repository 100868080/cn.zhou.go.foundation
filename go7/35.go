package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var area float64
	fmt.Println("输入a，b，c的值:")
	fmt.Scanf("%d %d %d", a, b, c)
	area = f1(a, b, c)
	fmt.Println(area)
}
func f1(a, b, c float64) float64 {
	var s, l float64
	s = (a + b + c) / 2
	l = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return l
}
