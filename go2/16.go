//位运算
package main

import (
	"fmt"
)

func main() {
	var a, b byte = 6, 11
	fmt.Println(a << 2)
	fmt.Println(b >> 2)
	fmt.Println(a &^ b)
}
