package main

import (
	"fmt"
)

func main() {
	type sdf struct {
		d string
	}
	var cf sdf
	cf = sdf{"5445"}
	fmt.Println(cf)
}
