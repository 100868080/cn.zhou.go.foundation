//指针变量的定义和引用
package main

import (
	"fmt"
)

func main() {
	var i int
	var i_pointer *int
	i = 100
	i_pointer = &i
	fmt.Println(*i_pointer)
}
