package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scanln(&a)

	switch {
	case (a > 2):
		fmt.Println("go")
		fallthrough
	case (a > 5):
		fmt.Println("lang")
	default:
		fmt.Println("c language")
	}
}
