package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)
	if a/3 == 2 {
		if a/3 == 2 {
			if a/6 == 1 {
				fmt.Println("go")
			}
		}
	} else {
		fmt.Println("lang")
	}
}
