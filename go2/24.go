package main

import (
	"fmt"
	"time"
)

func main() {
	var i float32
	i = 66.25555
	b := int(i)
	var m = string(b + 6)
	y := 32 + 112i
	fmt.Println(real(y))
	fmt.Println(imag(y))
	t := time.Now()
	e := time.Now()
	fmt.Println(e)
	fmt.Println(t)
	fmt.Println(i)
	fmt.Println(b)
	fmt.Println(m)
}
