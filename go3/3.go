package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c string = "false", "3.14", "1234"
	var x bool
	var y float64
	var z int64
	x, _ = strconv.ParseBool(a)
	y, _ = strconv.ParseFloat(b, 64)
	z, _ = strconv.ParseInt(c, 10, 0)
	fmt.Println(a, x)
	fmt.Println(b, y)
	fmt.Println(c, z)

}
