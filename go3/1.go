//字符串追加和替换
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "bnanafghkj"
	var count int = 3
	fmt.Println("ko" + strings.Repeat(s, count))
	fmt.Println("ko" + strings.Repeat(s, 2) + "logo")
	fmt.Println(strings.Replace("bna", "bna", "golang", 5))
	fmt.Println(strings.Replace(s, "na", "d", -1))
}
