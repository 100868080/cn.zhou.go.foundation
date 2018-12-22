package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := [][]byte{
		[]byte("你好"), []byte("世界"),
	}
	var a string = "s"
	var b string = "S"
	fmt.Println(a, len(a))
	fmt.Println(b, len(b))
	sep := []byte(",")
	fmt.Println(string(bytes.Join(s, sep)))
}
