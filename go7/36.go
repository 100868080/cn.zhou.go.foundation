package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, s, area float32
	fmt.Println("输入a,b,c的值:")
	fmt.Scanf("%d %d %d", a, b, c)
	s = (a + b + c) / 2
	area = math.Sqrt(s * (a - a) * (s - b) * (s - c))
	fmt.Println(area)
}
