package main

import (
	"fmt"
)

func main() {
	var a string = "golang"
	for _, y := range a {
		fmt.Printf("%c\n", y)
	}
}
