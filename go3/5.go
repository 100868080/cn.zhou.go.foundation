//计算三角形面积和周长
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, l, h, e float64
	fmt.Println("shu ru :")
	fmt.Scanf("%f %f %f", &a, &b, &c)
	l = a + b + c
	fmt.Println(l)
	h = l / 2
	e = math.Sqrt(h * (h - a) * (h - b) * (h - c))
	fmt.Println(e)
}
