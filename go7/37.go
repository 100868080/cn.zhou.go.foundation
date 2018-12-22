package main

import (
	"fmt"
)

func main() {

	var s = []float64{2.0, 2.3, 3.1, 3.25, 3.6, 6, 988}
	f1(s)

}
func f1(s []float64) {
	var r, g float64
	var a int
	for _, v := range s {
		r += v
	}
	a = len(s)

	g = r / float64(a)
	fmt.Println(g)
	return
}
