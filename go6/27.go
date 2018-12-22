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
	x := 0
	var _, y int
	for f, y := range week {
		x += y
	}

	fmt.Println(x)
}
