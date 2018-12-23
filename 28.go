package main

import (
	"fmt"
)

func main() {
	x := "Hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) //"HELLO (one letter per iteration)
	}
}
