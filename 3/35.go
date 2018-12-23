package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 167777216
	fmt.Println(f == f+1)
	const e = 2.71828
	fmt.Println(e)
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)

}

func compute() (value float64, ok bool) {
	if failed {
		return 0, false
	}
	return result, true
}
