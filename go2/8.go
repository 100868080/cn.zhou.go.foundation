//字节型数据的定义与使用
package main

import (
	"fmt"
)

func main() {
	var ch1, ch2 byte
	ch1 = 65
	ch2 = 'A'
	fmt.Println(string(ch1)) //要字符型输出需先转换
	fmt.Println(ch2)
}
