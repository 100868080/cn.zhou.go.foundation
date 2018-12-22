package main

import (
	"fmt"
)

func max(a, b int) (int, int) {
	/*if a > b {
		return a
	}*/
	return a * b, a + b
}
func myfunc(arg ...int) {

}

func main() {
	x := 3
	y := 4
	z := 6
	max_xy, max_x := max(x, y)
	max_xz, max_y := max(x, z)
	fmt.Println(max_xy, max_xz, max_x, max_y)

	/*for _, n := range arg {
		fmt.Printf("And the number is:%d\n", n)
	}*/

}
