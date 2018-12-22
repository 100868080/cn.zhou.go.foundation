//计算一元二次方程
package main

import (
	"fmt"
	"math"
)

func main() {
	var b, a, c, p, q, x, y, k float64
	fmt.Println("3 ge shu:")
	fmt.Scanf("%f %f %f", &a, &b, &c)
	k = b*b - 4*a*c
	p = -b / (2 * a)
	q = math.Sqrt(k) / 2 * a
	x = p + q
	y = p - q
	fmt.Printf("%6f\n", x)
	fmt.Printf("%6f\n", y)
}
