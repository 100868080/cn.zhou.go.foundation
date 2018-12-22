package main

import (
	"fmt"
)

func main() {
	var i_pointer *int
	var b int
	b = 1321564
	i_pointer = &b
	fmt.Println(*i_pointer)

}
