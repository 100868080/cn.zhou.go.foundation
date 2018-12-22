package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, b, c string = "go", "GO", "lang"
	fmt.Println(a, b, strings.EqualFold(a, b))
	fmt.Println(a, c, strings.EqualFold(b, c))
}
