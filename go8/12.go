package main

import (
	"fmt"
)

func main() {
	type f struct {
		dc int
	}
	type d struct {
		dc int
	}
	type dnf struct {
		d
		f
		//dc int
	}
	t := dnf{d{465}, f{456}}
	fmt.Println(t.f.dc, t.f.dc)
}
