package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Println("输入正整数a,b:")
	fmt.Scanf("%d,%d", &a, &b)
	c = mul(a, b)
	fmt.Printf("%d和%d最小公倍数是:%d\n", a, b, c)
}
func mul(a, b int) int {
	var d int
	d = div(a, b)
	return (a * b / d)
}
func div(a, b int) int {
	var r int
	for {
		r = a % b
		if r != 0 {
			a = b
			b = r
		} else {
			break
		}
	}
	return b
}
