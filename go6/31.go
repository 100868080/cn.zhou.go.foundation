package main

import (
	"fmt"
)

func main() {
	var a = []float64{3.14, 3.21, 3.22}
	var x float64
	for _, n := range a {
		x += n
	}
	fmt.Println(x / 3)
}
