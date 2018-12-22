package main

import (
	"fmt"
	"mymath"
)

func main() {
	var a, b int = 2, 3
	add := mymath.Add(a, b)
	fmt.Println(add)

}
