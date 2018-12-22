//使用变参求最小数
package main

import (
	"fmt"
)

func main() {
	min := mn(12, 3, 24, 7)
	fmt.Println("最小数为：", min)
	a := []int{-6, 13, 7, 25, -13}
	min = mn(a...)
	fmt.Println("数组最小数是:", min)
}
func mn(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
