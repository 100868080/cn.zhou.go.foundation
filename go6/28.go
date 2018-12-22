package main

import (
	"fmt"
)

func main() {
	week := map[string]int{
		"monday": 1, "tueafay": 2, "wednsday": 3,
		"thurday": 4, "friday": 5, "staurday": 6,
		"sunday": 7,
	}
	/*var y int
	var x string*/
	for x, y := range week {
		fmt.Println(x, y)
	}

	for o, k := range "go lang" {
		fmt.Println(o, k)

	}
	var g rune = 'l'
	fmt.Println(g)
}
