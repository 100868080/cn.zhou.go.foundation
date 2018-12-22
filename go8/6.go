package main

import (
	"fmt"
)

func main() {
	ma := map[string]struct {
		name string
		age  int
	}{
		/* "tacher":  */ "_": {"zhengzhi", 79},
		"student":            {"liming", 12},
	}
	fmt.Println(ma)
}
