package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf := []byte("golang")
	a := bytes.NewBuffer(buf)
	fmt.Println(string(a.Next(5)))
	a.UnreadByte()
	fmt.Println(string(a.Bytes()))
}
