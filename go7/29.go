//defer语句定义函数延迟执行
package main

import (
	"fmt"
)

func main() {
	f1()
	f2()
	fmt.Println()
	fmt.Println(f3())
}
func f1() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func f2() {
	for i := 0; i < 4; i++ {
		defer fmt.Printf("%d", i)
	}
}
func f3() (i int) {
	defer func() {
		i++
	}()
	return 1
}
