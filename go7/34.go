//多返回值找出字节切片中的整形数
package main

import (
	"fmt"
)

func main() {
	a := []byte{'1', '2', '3', '4', 'x', '5', 'h', '6', 'd', '7', 'h', '8'}
	var index, x, start int
	fmt.Println("输入起始索引位置:")
	fmt.Scanf("%d", &start)
	for {
		index, x = nextint(a, start)
		fmt.Printf("找到%d,下一个从%d开始\n", x, index)
		if index != 0 && index < len(a) {
			start = index
			continue
		} else {
			break
		}
	}
}
func nextint(b []byte, i int) (index, x int) {
	var start bool
	if b[i] >= '0' && b[i] <= '9' {
		start = true
	} else {
		start = false
	}
	for ; i < len(b); i++ {
		if b[i] >= '0' && b[i] <= '9' {
			if start == false {
				index = i
				break
			} else {
				x = x*10 + int(b[i]) - '0'
			}
		} else if b[i] < '0' || b[i] > '9' {
			start = false

		}

	}
	return index, x

}
