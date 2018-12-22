package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("the gone,mongsa zhike!")
	a := bytes.NewBuffer(buf)
	var d byte = ','
	n, err := a.ReadBytes(d)
	fmt.Println(string(n), err)
	fmt.Println(string(a.Next(4)))
	fmt.Println(string(a.Next(4)))
}
