package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "na"
	var cout int = 2
	fmt.Println("hello" + strings.Repeat(s, 1) + "wang")
	fmt.Println("golang" + strings.Repeat(s, cout))
	fmt.Println(strings.Replace("f", "k", "g", 3))
}
