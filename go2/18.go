package main

import (
	"fmt"
)

func main() {
	var str string
	str = "Hello 时间！"
	fmt.Println("unicode:")
	for i, ch := range str {
		fmt.Printf("str[%d]=%v\n", i, ch)
	}
}
