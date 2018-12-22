package main

import (
	"fmt"
	"strconv"
)

func main() {
	var b bool = false
	var f float64 = 3.14
	var k int64 = 1234
	fmt.Println(b, strconv.FormatBool(b))
	fmt.Println(f, strconv.FormatFloat(f, 'e', 3, 64))
	fmt.Println(k, strconv.FormatInt(k, 9))
}
