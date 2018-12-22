package main

import (
	"fmt"
)

func main() {
	m := map[string]struct {
		id   int
		name string
	}{
		"liming":  {321, "ming"},
		"student": {654, "ling"},
	}
	fmt.Println(m)
}
