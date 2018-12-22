package main

import (
	"fmt"
)

func main() {
	var map1 = map[string]int{
		"psw1": 123, "psw2": 456, "psw3": 789,
	}
	fmt.Println(map1)
	delete(map1, "psw2")
	fmt.Println(map1)
	map1["psw4"] = 135
	fmt.Println(map1)
}
