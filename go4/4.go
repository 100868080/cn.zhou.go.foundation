package main

import (
	"fmt"
)

func main() {
	var a, b int
	for a = 1; a <= 9; a++ {
		for b = 1; b <= a; b++ {
			fmt.Printf("%-2d *%-2d=%-7d", a, b, a*b)
			//fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

}
