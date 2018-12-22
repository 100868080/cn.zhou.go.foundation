package main

import (
	"fmt"
)

func main() {
	/*var a = map[string]int{}
	a["h"] = 3
	a["k"] = 65465
	//fmt.Println(a)
	fmt.Println(a)
	fmt.Println(a)*/

	/*var z map[string]int
	z = make(map[string]int)
	z["k1"] = 2
	fmt.Println(z)*/

	var k = map[string]int{
		"a": 1, "b": 2, "c": 3,
		"d": 4, "e": 5, "f": 6,
	}
	fmt.Println((k["a"] * k["b"]) % k["f"] * k["d"])
}
