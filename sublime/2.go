package main

import (
	"fmt"
)

func main() {
	s := "the sky"
	for v, k := range s {
		fmt.Println(v, k)
	}

	f := [8]int{1, 23, 3, 1, 3, 31, 2, 3}
	for h, j := range f {
		fmt.Println(h, j)
		for i := 0; i < 7; i++ {
		jk:
			int = jk + f[i]
			fmt.Println(jk)
		}
	}
}
