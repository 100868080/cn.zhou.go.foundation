package main

import (
	"fmt"
)

func main() {

	var re int = 0
	for i := 0; i < 20; i++ {
		re = fi(i)
		if i%5 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%4d", re)
	}
}

func fi(n int) (re int) {
	if n <= 1 {
		re = 1
	} else {
		re = fi(n-1) + fi(n-2)
	}
	return

}
