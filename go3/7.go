package main

import (
	"fmt"
	"math"
)
type pi float32
func main() {
	var r, x, y, k, float64
	fmt.Println("shu ru:")
	fmt.Scanf("%f,%f", &r, &k)
	x = r * math.Cos(k)
	y = r * math.Sin(k)
	fmt.Printf("%f\n", x)
	fmt.Printf("%f\n", y)
}
