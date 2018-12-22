package main

import (
	"fmt"
)

func main() {
	n := map[string]int{
		"git": 645,
		"pk":  466,
	}
	fmt.Println(n)
	//n =/* map[string]int*/{"pink": 645}
	delete(n, "git")
	n["note"] = 546
	n["pink"] = 9789898
	fmt.Println(n)
}
